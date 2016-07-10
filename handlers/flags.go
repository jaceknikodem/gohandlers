// Displays all flags within the binary.
//
// Usage:
//   http.Handle("/flags", *handlers.NewFlagHandler())
package handlers

import (
	"flag"
	"net/http"
)

// FlagInfo is an external structure exposed to consumers (template, JSON).
type FlagInfo struct {
	Flags map[string]string `json:"flags"`
}

type flagHandler struct {
}

// Expose implements Exposer interface.
func (h flagHandler) Expose(r *http.Request) (interface{}, error) {
	info := FlagInfo{
		Flags: make(map[string]string),
	}
	flag.VisitAll(func(f *flag.Flag) {
		info.Flags[f.Name] = f.Value.String()
	})
	return info, nil
}

func NewFlagHandler() *flagHandler {
	return &flagHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h flagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "flags.html")
}
