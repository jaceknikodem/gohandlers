package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jaceknikodem/gohandlers/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCounterOps(t *testing.T) {
	h := handlers.NewCounterHandler()

	handlers.Counters.Get("foo/bar").Increment()
	handlers.Counters.Get("foo/baz").IncrementBy(3)
	handlers.Counters.Get("bar/baz").IncrementBy(5)

	r, _ := http.NewRequest("GET", "/", nil)
	d, err := h.Expose(r)
	assert.Nil(t, err)
	info := d.(handlers.CountInfo)

	assert.Contains(t, info.Counters, "foo/bar")
	assert.Contains(t, info.Counters, "foo/baz")
	assert.Contains(t, info.Counters, "bar/baz")

	assert.Equal(t, info.Counters["foo/bar"], uint64(1))
	assert.Equal(t, info.Counters["foo/baz"], uint64(3))
	assert.Equal(t, info.Counters["bar/baz"], uint64(5))

	r, _ = http.NewRequest("GET", "/?prefix=foo/", nil)
	d, err = h.Expose(r)
	assert.Nil(t, err)
	info = d.(handlers.CountInfo)

	assert.Contains(t, info.Counters, "foo/bar")
	assert.Contains(t, info.Counters, "foo/baz")
	assert.NotContains(t, info.Counters, "bar/baz")
}
