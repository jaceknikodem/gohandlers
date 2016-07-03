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
	StartTime time.Time
	Uptime    time.Duration
}

type StatusHandler struct {
	Status statusInfo
}

// Expose defines structure exposed to external consumers.
func (h StatusHandler) Expose(r *http.Request) interface{} {
	info := StatusInfo{StartTime: h.Status.StartTime}
	info.Uptime = time.Since(info.StartTime)
	return info
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{
		Status: statusInfo{
			StartTime: time.Now(),
		},
	}
}

func (h StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "status.html")
}
