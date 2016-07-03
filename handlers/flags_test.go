package handlers

import (
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	_ = flag.Int("fake", 123, "Fake flag for testing.")

	h := NewFlagHandler()

	r, _ := http.NewRequest("GET", "/", nil)
	data := h.Expose(r)
	info := data.(FlagInfo)

	assert.Contains(t, info.Flags, "fake")
	assert.Equal(t, info.Flags["fake"], "123")
}
