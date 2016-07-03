package handlers

import "net/http"

const allKey = "all"

// responseWriter is just a thin wrapper around http.ResponseWriter that counts response size
type responseWriter struct {
	writer http.ResponseWriter
	size   int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		writer: w,
	}
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.writer.Write(b)
	w.size += size
	return size, err
}

func (w *responseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *responseWriter) WriteHeader(code int) {
	w.writer.WriteHeader(code)
}

func (w responseWriter) Size() int {
	return w.size
}

// RequestMiddleware counts request/response size and numbers off calls
type RequestMiddleware struct {
	Calls        *CounterSet
	RequestSize  *CounterSet
	ResponseSize *CounterSet
}

func NewRequestMiddleware() *RequestMiddleware {
	return &RequestMiddleware{
		Calls:        NewCounterSet(),
		RequestSize:  NewCounterSet(),
		ResponseSize: NewCounterSet(),
	}
}

func (m *RequestMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := newResponseWriter(w)

		m.Calls.Get(allKey).Increment()
		m.RequestSize.Get(allKey).IncrementBy(uint64(r.ContentLength))

		next.ServeHTTP(nw, r)

		m.ResponseSize.Get(allKey).IncrementBy(uint64(nw.Size()))
	})
}
