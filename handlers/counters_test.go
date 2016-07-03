package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterOps(t *testing.T) {
	cs := NewCounterSet()
	cs.Get("foo/bar").Increment()
	v := cs.Get("foo/bar").Get()
	assert.Equal(t, v, uint64(1))
}
