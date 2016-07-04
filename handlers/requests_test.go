package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaceknikodem/gohandlers/handlers"
	"github.com/stretchr/testify/assert"
)

type fakeHandler struct{}

func (h fakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func TestRequestHandler(t *testing.T) {
	m := handlers.NewRequestMiddleware()
	h := m.Wrap(fakeHandler{})

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	assert.Equal(t, m.Calls.Get("all").Value(), uint64(1))
	assert.Equal(t, m.RequestSize.Get("all").Value(), uint64(0))
	assert.Equal(t, m.ResponseSize.Get("all").Value(), uint64(4))
}
