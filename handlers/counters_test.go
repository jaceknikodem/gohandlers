package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterOps(t *testing.T) {
	h := NewCounterHandler()

	Counters.Get("foo/bar").Increment()
	Counters.Get("foo/baz").IncrementBy(3)
	Counters.Get("bar/baz").IncrementBy(5)

	r, _ := http.NewRequest("GET", "/", nil)
	d := h.Expose(r)
	info := d.(countInfo)

	assert.Contains(t, info.Counters, "foo/bar")
	assert.Contains(t, info.Counters, "foo/baz")
	assert.Contains(t, info.Counters, "bar/baz")

	assert.Equal(t, info.Counters["foo/bar"], uint64(1))
	assert.Equal(t, info.Counters["foo/baz"], uint64(3))
	assert.Equal(t, info.Counters["bar/baz"], uint64(5))

	r, _ = http.NewRequest("GET", "/?prefix=foo/", nil)
	d = h.Expose(r)
	info = d.(countInfo)

	assert.Contains(t, info.Counters, "foo/bar")
	assert.Contains(t, info.Counters, "foo/baz")
	assert.NotContains(t, info.Counters, "bar/baz")
}
