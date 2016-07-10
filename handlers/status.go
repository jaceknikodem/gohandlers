// The status page displays:
// * start time and uptime
// * links to other subpages
//
// Example usages:
//   http.Handle("/status", *handlers.NewStatusHandler())
package handlers

import (
	"net"
	"net/http"
	"time"
)

func addresses() []string {
	ifaces, _ := net.Interfaces()
	rv := make([]string, 0, len(ifaces))
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return rv
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					rv = append(rv, ipnet.IP.String())
				}
			}
		}
	}
	return rv
}

// statusInfo is an internal structure used to gather all information
type statusInfo struct {
	StartTime time.Time
}

// StatusInfo is an external structure exposed to consumers (template, JSON)
type StatusInfo struct {
	StartTime time.Time     `json:"start_time"`
	Uptime    time.Duration `json:"uptime"`
	Addresses []string      `json:"addresses"`
}

type statusHandler struct {
	Status statusInfo
}

// Expose defines structure exposed to external consumers.
func (h statusHandler) Expose(r *http.Request) (interface{}, error) {
	info := StatusInfo{
		StartTime: h.Status.StartTime,
		Uptime:    time.Since(h.Status.StartTime),
		Addresses: addresses(),
	}
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
