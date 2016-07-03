package handlers

import (
	"flag"
	"net/http"
)

type FlagInfo struct {
	Flags map[string]string
}

type FlagHandler struct {
}

func (h FlagHandler) Expose() interface{} {
	info := FlagInfo{
		Flags: make(map[string]string),
	}
	flag.VisitAll(func(f *flag.Flag) {
		info.Flags[f.Name] = f.Value.String()
	})
	return info
}

func NewFlagHandler() *FlagHandler {
	return &FlagHandler{}
}

func (h FlagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "flags.html")
}