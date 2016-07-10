// Display all environment variables.
//
// Usage:
//   http.Handle("/env", *handlers.NewEnvHandler())
package handlers

import (
	"net/http"
	"os"
	"strings"
)

// EnvInfo is a struct exposed externally.
type EnvInfo struct {
	Vars map[string]string `json:"vars"`
}

type envHandler struct {
	Info EnvInfo
}

// Expose implements Exposer interface.
func (h envHandler) Expose(r *http.Request) (interface{}, error) {
	return h.Info, nil
}

// NewEnvHandler creates a new envHandler.
func NewEnvHandler() *envHandler {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		p := strings.SplitN(e, "=", 2)
		vars[p[0]] = p[1]
	}
	return &envHandler{
		Info: EnvInfo{
			Vars: vars,
		},
	}
}

// ServeHTTP implements http.Handler interface.
func (h envHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "env.html")
}
