package handlers

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	k := "TEST_FAKE_ENV_VAR"
	err := os.Setenv(k, "fake_value")
	assert.Nil(t, err)
	defer os.Unsetenv(k)

	h := NewEnvHandler()

	r, _ := http.NewRequest("GET", "/", nil)
	d, err := h.Expose(r)
	assert.Nil(t, err)
	info := d.(EnvInfo)

	assert.Contains(t, info.Vars, k)
}
