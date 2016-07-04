package handlers_test

import (
	"flag"
	"net/http"
	"testing"

	"github.com/jaceknikodem/gohandlers/handlers"
	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	_ = flag.Int("fake", 123, "Fake flag for testing.")

	h := handlers.NewFlagHandler()

	r, _ := http.NewRequest("GET", "/", nil)
	data := h.Expose(r)
	info := data.(handlers.FlagInfo)

	assert.Contains(t, info.Flags, "fake")
	assert.Equal(t, info.Flags["fake"], "123")
}
