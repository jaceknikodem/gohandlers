package handlers

import (
	"net/http"
	"os"
	"strings"
)

type EnvInfo struct {
	Vars map[string]string
}

type EnvHandler struct {
	Info EnvInfo
}

func (h EnvHandler) Expose() interface{} {
	return h.Info
}

func NewEnvHandler() *EnvHandler {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		p := strings.SplitN(e, "=", 2)
		vars[p[0]] = p[1]
	}
	return &EnvHandler{
		Info: EnvInfo{
			Vars: vars,
		},
	}
}

func (h EnvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "env.html")
}
