package handlers_test

import (
	"flag"
	"net/http"
	"testing"

	"github.com/jaceknikodem/gohandlers/handlers"
	"github.com/stretchr/testify/assert"
)

func assertFlagEqual(t *testing.T, info handlers.FlagInfo, name string, expected string) {
	assert.Contains(t, info.Flags, name)
	assert.Equal(t, info.Flags[name], expected)
}

func TestFlag(t *testing.T) {
	_ = flag.Int("fake", 123, "Fake flag for testing.")

	h := handlers.NewFlagHandler()
	r, _ := http.NewRequest("GET", "/", nil)

	data, err := h.Expose(r)
	assert.NoError(t, err)
	info := data.(handlers.FlagInfo)
	assertFlagEqual(t, info, "fake", "123")
}

func TestUpdateFlagSuccessful(t *testing.T) {
	_ = flag.Int("foo", 123, "Fake flag for testing.")

	h := handlers.NewFlagHandler()
	r, _ := http.NewRequest("GET", "/?name=foo&value=234", nil)

	data, err := h.Expose(r)
	assert.NoError(t, err)
	info := data.(handlers.FlagInfo)
	assertFlagEqual(t, info, "foo", "234")
}

func TestUpdateFlagFailNonExistingFlag(t *testing.T) {
	_ = flag.Int("bar", 123, "Fake flag for testing.")

	h := handlers.NewFlagHandler()
	r, _ := http.NewRequest("GET", "/?name=nonexisting&value=234", nil)

	_, err := h.Expose(r)
	assert.Error(t, err)
}

func TestUpdateFlagFailBadType(t *testing.T) {
	_ = flag.Int("baz", 123, "Fake flag for testing.")

	h := handlers.NewFlagHandler()
	r, _ := http.NewRequest("GET", "/?name=baz&value=wrongtype", nil)

	_, err := h.Expose(r)
	assert.Error(t, err)

	r, _ = http.NewRequest("GET", "/", nil)
	data, err := h.Expose(r)
	assert.NoError(t, err)
	info := data.(handlers.FlagInfo)
	assertFlagEqual(t, info, "baz", "123")
}
