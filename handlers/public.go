package handlers

import "net/http"

var reqMid = NewRequestMiddleware()

func RegisterAll() {
	http.Handle("/kill", *NewKillHandler())
	http.Handle("/status", *NewStatusHandler())
	http.Handle("/env", *NewEnvHandler())
	http.Handle("/counts", *NewCounterHandler())
	http.Handle("/flags", *NewFlagHandler())
	http.Handle("/requests", *reqMid)
}

func Handle(pattern string, h http.Handler) {
	http.Handle(pattern, reqMid.Wrap(h))
}
