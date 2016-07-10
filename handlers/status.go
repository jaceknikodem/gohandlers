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
	"os"
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
	Hostname  string
	Pid       int
}

// StatusInfo converts statusInfo (the internal struct) to StatusInfo (the public struct).
func (s statusInfo) StatusInfo() StatusInfo {
	return StatusInfo{
		StartTime: s.StartTime,
		Hostname:  s.Hostname,
		Pid:       s.Pid,
	}
}

// StatusInfo is an external structure exposed to consumers (template, JSON)
type StatusInfo struct {
	StartTime time.Time     `json:"start_time"`
	Uptime    time.Duration `json:"uptime"`
	Addresses []string      `json:"addresses"`
	Cwd       string        `json:"cwd"`
	Hostname  string        `json:"hostname"`
	Pid       int           `json:"pid"`
}

type statusHandler struct {
	Status statusInfo
}

// Expose defines structure exposed to external consumers.
func (h statusHandler) Expose(r *http.Request) (interface{}, error) {
	info := h.Status.StatusInfo()
	info.Uptime = time.Since(info.StartTime)
	info.Addresses = addresses()
	if d, e := os.Getwd(); e == nil {
		info.Cwd = d
	}
	return info, nil
}

// NewStatusHandler creates a new statusHandler.
func NewStatusHandler() *statusHandler {
	s := statusInfo{
		StartTime: time.Now(),
		Pid:       os.Getegid(),
	}
	if n, e := os.Hostname(); e == nil {
		s.Hostname = n
	}
	return &statusHandler{
		Status: s,
	}
}

// ServeHTTP implements http.Handler interface.
func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "status.html")
}
