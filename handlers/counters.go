package handlers

import (
	"fmt"
	"net/http"
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

func (c *Counter) Get() uint64 {
	return atomic.LoadUint64(&c.value)
}

func (c Counter) String() string {
	return fmt.Sprintf("%s = %d", c.name, c.Get())
}

type CounterSet struct {
	counters map[string]*Counter
}

func NewCounterSet() *CounterSet {
	return &CounterSet{
		counters: make(map[string]*Counter),
	}
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

type CountInfo struct {
	Counters map[string]uint64
}

type CounterHandler struct{}

func (h CounterHandler) Expose() interface{} {
	info := CountInfo{
		Counters: make(map[string]uint64),
	}
	for name, c := range Counters.counters {
		info.Counters[name] = c.Get()
	}
	return info
}

func NewCounterHandler() *CounterHandler {
	return &CounterHandler{}
}

func (h CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveHTTP(w, r, h, "counters.html")
}
