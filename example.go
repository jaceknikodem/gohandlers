// Binary example shows how to attach handlers to a server.
// * status handler
// * counter handler
// * request handler
// * variables handler
// * flags handler
package main

import (
	"fmt"
	"net/http"

	"github.com/jaceknikodem/gohandlers/handlers"
)

func main() {
	fmt.Println("Starting a server")

	http.Handle("/status", *handlers.NewStatusHandler())
	http.Handle("/env", *handlers.NewEnvHandler())
	http.Handle("/counts", *handlers.NewCounterHandler())

	handlers.Counters.Get("foo/bar").IncrementBy(5)

	http.ListenAndServe(":8080", nil)
}
