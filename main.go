package main

import "fmt"

func setupHandlers(app *App) {
	app.HandleFunc("/view", ViewURL)
	app.HandleFunc("/add", AddURL)
	app.HandleFunc("/r/", GotoURL)
	//app.Mux.Handle("/static/", http.FileServer(http.Dir("./static")))
}

func main() {
	app := NewApp()
	setupHandlers(app)
	if err := app.Start(); err != nil {
		fmt.Println(err)
	}
}
