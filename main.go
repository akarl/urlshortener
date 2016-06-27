package main

func setupHandlers(app *App) {
	app.HandleFunc("/view", ViewURL)
	app.HandleFunc("/add", AddURL)
	app.HandleFunc("/r/", GotoURL)
}

func main() {
	app := NewApp()
	setupHandlers(app)
	app.Start()
}
