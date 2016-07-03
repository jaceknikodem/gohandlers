package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
)

var Counters CounterSet

func init() {
	Counters = *NewCounterSet()
}

type Counter struct {
	name  string
	value uint64
}

func (c *Counter) Increment() {
	atomic.AddUint64(&(c.value), 1)
}

func (c *Counter) IncrementBy(n uint64) {
	atomic.AddUint64(&(c.value), n)
}

func (c *Counter) Value() uint64 {
	return atomic.LoadUint64(&c.value)
}

func (c Counter) String() string {
	return fmt.Sprintf("%s = %d", c.name, c.Value())
}

type CounterSet struct {
	counters map[string]*Counter
}

func NewCounterSet() *CounterSet {
	return &CounterSet{
		counters: make(map[string]*Counter),
	}
}

func (cs CounterSet) String() string {
	var buffer bytes.Buffer
	for _, c := range cs.counters {
		buffer.WriteString(fmt.Sprintf("%s\n", c))
	}
	return buffer.String()
}

func (cs *CounterSet) Get(name string) *Counter {
	if c, ok := cs.counters[name]; ok {
		return c
	}
	c := new(Counter)
	c.name = name
	cs.counters[name] = c
	return cs.counters[name]
}

func (cs CounterSet) WithPrefix(p string) *CounterSet {
	ncs := NewCounterSet()
	for n, c := range cs.counters {
		if strings.HasPrefix(n, p) {
			ncs.counters[n] = c
		}
	}
	return ncs
}

type CountInfo struct {
	Counters map[string]uint64
}

type CounterHandler struct{}

func (h CounterHandler) Expose(r *http.Request) interface{} {
	info := CountInfo{
		Counters: make(map[string]uint64),
	}
	q := r.URL.Query()
	prefix, _ := getValue(&q, "prefix")
	for name, c := range Counters.counters {
		if strings.HasPrefix(name, prefix) {
			info.Counters[name] = c.Value()
		}
	}
	return info
}

func NewCounterHandler() *CounterHandler {
	return &CounterHandler{}
}

func (h CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "counters.html")
}
