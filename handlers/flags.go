// Displays all flags within the binary.
//
// Usage:
//   http.Handle("/flags", *handlers.NewFlagHandler())
package handlers

import (
	"flag"
	"net/http"
	"net/url"
)

// FlagInfo is an external structure exposed to consumers (template, JSON).
type FlagInfo struct {
	Flags map[string]string `json:"flags"`
}

type flagHandler struct {
}

func updateFlag(vs url.Values) error {
	values, err := GetValues(vs, "name", "value")
	if err != nil {
		// Updating is optional, so no error here.
		return nil
	}
	name, value := values[0], values[1]
	if err := flag.CommandLine.Set(name, value); err != nil {
		return err
	}
	return nil
}

// Expose implements Exposer interface.
func (h flagHandler) Expose(r *http.Request) (interface{}, error) {
	if err := updateFlag(r.URL.Query()); err != nil {
		return nil, err
	}

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
