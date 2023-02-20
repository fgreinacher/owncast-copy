package admin

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetHardwareStats will return hardware utilization over time.
func (c *Controller) GetHardwareStats(w http.ResponseWriter, r *http.Request) {
	m := c.Service.Metrics.GetMetrics()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Errorln(err)
	}
}
