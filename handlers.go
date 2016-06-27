package main

import (
	"log"
	"net/http"
)

func GotoURL(w http.ResponseWriter, r *http.Request, app *App) {
	code := r.URL.Path[len("/r/"):]
	url, err := app.URLStorage.GetURL(code)
	if err != nil {
		log.Panic(err)
	}
	if url == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func AddURL(w http.ResponseWriter, r *http.Request, app *App) {
	if r.Method == http.MethodPost {
		url := r.FormValue("url")

		if code, err := app.URLStorage.AddURL(url); err == nil {
			http.Redirect(w, r, "/view?c="+code, http.StatusFound)
		} else {
			log.Panic(err)
		}
	} else {
		if err := app.Templates.ExecuteTemplate(w, "add.html", nil); err != nil {
			log.Panic(err)
		}
	}
}

func ViewURL(w http.ResponseWriter, r *http.Request, app *App) {
	code := r.FormValue("c")
	if code == "" {
		http.NotFound(w, r)
		return
	}

	url, err := app.URLStorage.GetURL(code)
	if err != nil {
		log.Panic(err)
	}
	if url == "" {
		http.Redirect(w, r, "/add", http.StatusFound)
		return
	}

	data := struct{ Code, URL string }{code, url}
	if err := app.Templates.ExecuteTemplate(w, "view.html", data); err != nil {
		log.Panic(err)
	}
}
