package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var templates *template.Template

func registerHandlers(router *httprouter.Router) {
	router.GET("/", listUrls)
	router.POST("/create", createURL)
	router.GET("/delete/:code", deleteURL)
	router.GET("/r/:code", redirectToURL)
}

func registerTemplates() {
	templates = template.Must(template.ParseFiles(
		"templates/base.tmpl",
		"templates/list_urls.tmpl",
	))
}

func StartHTTPServer() error {
	router := httprouter.New()
	registerTemplates()
	registerHandlers(router)

	return http.ListenAndServe("localhost:8080", router)
}
