// Binary example shows how to attach handlers to a server.
// * status handler
// * counter handler
// * request handler
// * variables handler
// * flags handler
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jaceknikodem/gohandlers/handlers"
)

var port = flag.Int("port", 8080, "Port to run on.")

type fakeHandler struct{}

func (h fakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "random content")
}

func main() {
	p := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting a server on %s\n", p)

	handlers.Handle("/", fakeHandler{})
	handlers.RegisterAll()

	handlers.Counters.Get("foo/bar").IncrementBy(5)

	http.ListenAndServe(p, nil)
}
