// Handler to kill the process immediately. No clean-up.
package handlers

import (
	"net/http"
	"os"
)

type killHandler struct {
}

func NewKillHandler() *killHandler {
	return &killHandler{}
}

func (h killHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os.Exit(1)
}
