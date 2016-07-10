package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var ErrMissingAttribute = errors.New("Attribute is missing")

// Exposer exposes an object that is an external representation visible to all consumers.
type Exposer interface {
	Expose(*http.Request) (interface{}, error)
}

func GetValue(values url.Values, key string) (string, bool) {
	if value, ok := values[key]; ok && len(value) > 0 {
		return value[0], true
	}
	return "", false
}

func GetValues(vs url.Values, names ...string) ([]string, error) {
	flags := make([]string, 0, len(names))
	for _, name := range names {
		if v, ok := GetValue(vs, name); ok {
			flags = append(flags, v)
		} else {
			return flags, ErrMissingAttribute
		}
	}
	return flags, nil
}

func jsonRequested(values url.Values) bool {
	if format, ok := GetValue(values, "format"); ok && format == "json" {
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
	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func renderError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(struct {
		Error string
	}{
		e.Error(),
	})
}

func serveHTTP(w http.ResponseWriter, r *http.Request, e Exposer, name string) {
	data, err := e.Expose(r)
	if err != nil {
		renderError(w, err)
		return
	}

	if jsonRequested(r.URL.Query()) {
		json.NewEncoder(w).Encode(data)
		return
	}

	renderTemplate(w, name, data)
}
