// Usage:
//   http.Handle("/counts", *handlers.NewCounterHandler())
package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
)

var Counters counterSet

func init() {
	Counters = *newCounterSet()
}

type counter struct {
	name  string
	value uint64
}

// Increment increments a counter by 1 atomically.
func (c *counter) Increment() {
	atomic.AddUint64(&(c.value), 1)
}

// IncrementBy increments a counter by a given number atomically.
func (c *counter) IncrementBy(n uint64) {
	atomic.AddUint64(&(c.value), n)
}

// Value gets counter's value.
func (c *counter) Value() uint64 {
	return atomic.LoadUint64(&c.value)
}

// String implements fmt.Stringer interface.
func (c counter) String() string {
	return fmt.Sprintf("%s = %d", c.name, c.Value())
}

type counterSet struct {
	counters map[string]*counter
}

func newCounterSet() *counterSet {
	return &counterSet{
		counters: make(map[string]*counter),
	}
}

// String implements fmt.Stringer interface.
func (cs counterSet) String() string {
	var buffer bytes.Buffer
	for _, c := range cs.counters {
		buffer.WriteString(fmt.Sprintf("%s\n", c))
	}
	return buffer.String()
}

// Get looks up or creates a new counter.
func (cs *counterSet) Get(name string) *counter {
	if c, ok := cs.counters[name]; ok {
		return c
	}
	c := new(counter)
	c.name = name
	cs.counters[name] = c
	return cs.counters[name]
}

// WithPrefix returns a counterSet with counters starting with a given string.
func (cs counterSet) WithPrefix(p string) *counterSet {
	ncs := newCounterSet()
	for n, c := range cs.counters {
		if strings.HasPrefix(n, p) {
			ncs.counters[n] = c
		}
	}
	return ncs
}

// CountInfo returns counterSet represented as a CountInfo,
func (cs counterSet) CountInfo() CountInfo {
	info := CountInfo{
		Counters: make(map[string]uint64),
	}
	for name, c := range cs.counters {
		info.Counters[name] = c.Value()
	}
	return info
}

// CountInfo is an external structure exposed to consumers (template, JSON).
type CountInfo struct {
	Counters map[string]uint64 `json:"counters"`
}

type counterHandler struct{}

// Expose implements Exposer interface.
func (h counterHandler) Expose(r *http.Request) interface{} {
	q := r.URL.Query()
	prefix, _ := GetValue(&q, "prefix")
	cs := Counters.WithPrefix(prefix)
	return cs.CountInfo()
}

// NewCounterHandler creates a new counterHandler.
func NewCounterHandler() *counterHandler {
	return &counterHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "counters.html")
}
