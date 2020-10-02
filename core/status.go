package core

import (
	"time"

	"github.com/gabek/owncast/config"
	"github.com/gabek/owncast/core/ffmpeg"
	"github.com/gabek/owncast/models"
	"github.com/gabek/owncast/utils"
)

//GetStatus gets the status of the system
func GetStatus() models.Status {
	if _stats == nil {
		return models.Status{}
	}

	return models.Status{
		Online:                IsStreamConnected(),
		ViewerCount:           len(_stats.Clients),
		OverallMaxViewerCount: _stats.OverallMaxViewerCount,
		SessionMaxViewerCount: _stats.SessionMaxViewerCount,
		LastDisconnectTime:    _stats.LastDisconnectTime,
		LastConnectTime:       _stats.LastConnectTime,
	}
}

//SetStreamAsConnected sets the stream as connected
func SetStreamAsConnected() {
	stopCleanupTimer()

	_stats.StreamConnected = true
	_stats.LastConnectTime = utils.NullTime{time.Now(), true}
	_stats.LastDisconnectTime = utils.NullTime{time.Now(), false}

	chunkPath := config.Config.GetPublicHLSSavePath()
	if usingExternalStorage {
		chunkPath = config.Config.GetPrivateHLSSavePath()
	}

	if _yp != nil {
		_yp.Start()
	}

	ffmpeg.StartThumbnailGenerator(chunkPath, config.Config.VideoSettings.HighestQualityStreamIndex)
}

//SetStreamAsDisconnected sets the stream as disconnected
func SetStreamAsDisconnected() {
	_stats.StreamConnected = false
	_stats.LastDisconnectTime = utils.NullTime{time.Now(), true}

	if _yp != nil {
		_yp.Stop()
	}

	ffmpeg.ShowStreamOfflineState()
	startCleanupTimer()
}
