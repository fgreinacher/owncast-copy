package controllers

import (
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/core"
	"github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/router/middleware"
	"github.com/owncast/owncast/utils"
	log "github.com/sirupsen/logrus"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
)

type FileServerHandler struct {
	HLSPath string
}

func (fsh *FileServerHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, fsh.HLSPath)
}

// HandleHLSRequest will manage all requests to HLS content.
func HandleHLSRequest(w http.ResponseWriter, r *http.Request) {
	// Sanity check to limit requests to HLS file types.
	if filepath.Ext(r.URL.Path) != ".m3u8" && filepath.Ext(r.URL.Path) != ".ts" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseCache, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(20),
		memory.AdapterWithStorageCapacity(209_715_200),
	)
	if err != nil {
		log.Warn("unable to create web cache", err)
	}

	// Since HLS segments cannot be changed once they're rendered, we can cache
	// individual segments for a long time.
	longTermHLSSegmentCache, err := cache.NewClient(
		cache.ClientWithAdapter(responseCache),
		cache.ClientWithTTL(30*time.Second),
		cache.ClientWithExpiresHeader(),
	)
	if err != nil {
		log.Warn("unable to create web cache client", err)
	}

	requestedPath := r.URL.Path
	relativePath := strings.Replace(requestedPath, "/hls/", "", 1)
	fullPath := filepath.Join(config.HLSStoragePath, relativePath)

	// If using external storage then only allow requests for the
	// master playlist at stream.m3u8, no variants or segments.
	if data.GetS3Config().Enabled && relativePath != "stream.m3u8" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Handle playlists
	if path.Ext(r.URL.Path) == ".m3u8" {
		// Playlists should never be cached.
		middleware.DisableCache(w)

		// Force the correct content type
		w.Header().Set("Content-Type", "application/x-mpegURL")

		// Use this as an opportunity to mark this viewer as active.
		viewer := models.GenerateViewerFromRequest(r)
		core.SetViewerActive(&viewer)
	} else {
		cacheTime := utils.GetCacheDurationSecondsForPath(relativePath)
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheTime))

		fileServer := &FileServerHandler{HLSPath: fullPath}
		longTermHLSSegmentCache.Middleware(fileServer).ServeHTTP(w, r)
		// http.ServeH
		// http.Handle("/hls", longTermHLSSegmentCache.Middleware(fileServer))
		return
	}

	middleware.EnableCors(w)
	http.ServeFile(w, r, fullPath)
}
