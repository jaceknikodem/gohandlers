// The status page displays:
// * start time and uptime
// * links to other subpages
//
// Example usages:
//   http.Handle("/status", *handlers.NewStatusHandler())
package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// statusInfo is an internal structure used to gather all information
type statusInfo struct {
	StartTime time.Time
}

// StatusInfo is an external structure exposed to consumers (template, JSON)
type StatusInfo struct {
	StartTime time.Time
	Uptime    time.Duration
}

// Expose defines structure exposed to external consumers.
func (i *statusInfo) Expose() StatusInfo {
	info := StatusInfo{StartTime: i.StartTime}
	info.Uptime = time.Since(info.StartTime)
	return info
}

type StatusHandler struct {
	Status statusInfo
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{
		Status: statusInfo{
			StartTime: time.Now(),
		},
	}
}

func (h StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := h.Status.Expose()

	query := r.URL.Query()
	if jsonRequested(&query) {
		json.NewEncoder(w).Encode(i)
		return
	}
	renderTemplate(w, i)
}
