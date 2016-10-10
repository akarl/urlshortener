package main

import (
	"html/template"
	"net/http"
	"time"
)

type HandleFunc func(http.ResponseWriter, *http.Request, *App)

type App struct {
	Server     *http.Server
	Mux        *http.ServeMux
	Templates  *template.Template
	URLStorage *URLStorage
}

func NewApp() *App {
	mux := http.NewServeMux()
	pool := NewRedisPool()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
	}

	st := &URLStorage{
		RedisPool: pool,
	}

	return &App{
		Server:     s,
		Mux:        mux,
		URLStorage: st,
	}
}

func (app *App) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"templates/base.tmpl",
		"templates/" + name,
	)

	if err != nil {
		http.Error(w, "", 500)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "", 500)
	}
}

func (app *App) HandleFunc(pattern string, handler HandleFunc) {
	app.Mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	})
}

func (app *App) Start() error {
	return app.Server.ListenAndServe()
}
