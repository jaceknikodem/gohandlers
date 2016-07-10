// The status page displays:
// * start time and uptime
// * links to other subpages
//
// Example usages:
//   http.Handle("/status", *handlers.NewStatusHandler())
package handlers

import (
	"net/http"
	"time"
)

// statusInfo is an internal structure used to gather all information
type statusInfo struct {
	StartTime time.Time
}

// StatusInfo is an external structure exposed to consumers (template, JSON)
type StatusInfo struct {
	StartTime time.Time     `json:"start_time"`
	Uptime    time.Duration `json:"uptime"`
}

type statusHandler struct {
	Status statusInfo
}

// Expose defines structure exposed to external consumers.
func (h statusHandler) Expose(r *http.Request) (interface{}, error) {
	info := StatusInfo{StartTime: h.Status.StartTime}
	info.Uptime = time.Since(info.StartTime)
	return info, nil
}

// NewStatusHandler creates a new statusHandler.
func NewStatusHandler() *statusHandler {
	return &statusHandler{
		Status: statusInfo{
			StartTime: time.Now(),
		},
	}
}

// ServeHTTP implements http.Handler interface.
func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "status.html")
}
