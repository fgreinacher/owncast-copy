package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/utils"
)

// ReportPlaybackMetrics will accept playback metrics from a client and save
// them for future video health reporting.
func (s *Service) ReportPlaybackMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		s.WriteSimpleResponse(w, false, r.Method+" not supported")
		return
	}

	type reportPlaybackMetricsRequest struct {
		Bandwidth             float64 `json:"bandwidth"`
		Latency               float64 `json:"latency"`
		Errors                float64 `json:"errors"`
		DownloadDuration      float64 `json:"downloadDuration"`
		QualityVariantChanges float64 `json:"qualityVariantChanges"`
	}

	decoder := json.NewDecoder(r.Body)
	var request reportPlaybackMetricsRequest
	if err := decoder.Decode(&request); err != nil {
		log.Errorln("error decoding playback metrics payload:", err)
		s.WriteSimpleResponse(w, false, err.Error())
		return
	}

	clientID := utils.GenerateClientIDFromRequest(r)

	s.Metrics.RegisterPlaybackErrorCount(clientID, request.Errors)
	if request.Bandwidth != 0.0 {
		s.Metrics.RegisterPlayerBandwidth(clientID, request.Bandwidth)
	}

	if request.Latency != 0.0 {
		s.Metrics.RegisterPlayerLatency(clientID, request.Latency)
	}

	if request.DownloadDuration != 0.0 {
		s.Metrics.RegisterPlayerSegmentDownloadDuration(clientID, request.DownloadDuration)
	}

	s.Metrics.RegisterQualityVariantChangesCount(clientID, request.QualityVariantChanges)
}
