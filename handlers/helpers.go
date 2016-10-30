package handlers

import "net/http"

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
