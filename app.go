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
		Templates:  template.Must(template.ParseFiles("add.html", "view.html")),
		URLStorage: st,
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
