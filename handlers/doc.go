// handlers exposes multiple handlers that add extra functionality to a server
// e.g. a status page, an RPC summary page, etc.
//
// Example usage:
//   handlers.Handle("/", fakeHandler{})
//   handlers.RegisterAll()
//   http.ListenAndServe(p, nil)
//
package handlers
