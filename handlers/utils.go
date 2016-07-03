package handlers

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

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

func renderTemplate(w http.ResponseWriter, data interface{}) {
	t := template.New("status.html")
	// TODO(jaceknikodem): Move template loading to init.
	t, err := t.ParseFiles("templates/status.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
