package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// Exposer exposes an object that is an external representation visible to all consumers.
type Exposer interface {
	Expose() interface{}
}

func getValue(values *url.Values, key string) (string, bool) {
	if value, ok := (*values)[key]; ok && len(value) > 0 {
		return value[0], true
	}
	return "", false
}

func jsonRequested(values *url.Values) bool {
	if format, ok := getValue(values, "format"); ok && format == "json" {
		return true
	}
	return false
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	t := template.New(name)
	// TODO(jaceknikodem): Move template loading to init.
	t, err := t.ParseFiles("templates/" + name)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func serveHTTP(w http.ResponseWriter, r *http.Request, e Exposer, name string) {
	ext := e.Expose()

	q := r.URL.Query()
	if jsonRequested(&q) {
		json.NewEncoder(w).Encode(ext)
		return
	}

	renderTemplate(w, name, ext)
}
