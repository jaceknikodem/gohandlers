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
	data, err := h.Expose(r)
	assert.Nil(t, err)
	info := data.(handlers.FlagInfo)

	assert.Contains(t, info.Flags, "fake")
	assert.Equal(t, info.Flags["fake"], "123")
}

// func TestUpdateFlag(t *testing.T) {
// 	_ = flag.Int("fake", 123, "Fake flag for testing.")

// 	h := handlers.NewFlagHandler()

// 	r, _ := http.NewRequest("GET", "/?name=fake&value=234", nil)

// }
