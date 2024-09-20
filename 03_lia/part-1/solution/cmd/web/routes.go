package main

import "net/http"

// routes returns a mux with all the registered routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /account/create", app.accountCreate)
	mux.HandleFunc("POST /account/create", app.accountCreatePost)
	mux.HandleFunc("GET /account/view/{id}", app.accountView)

	mux.HandleFunc("GET /thread/create", app.threadCreate)
	mux.HandleFunc("POST /thread/create", app.threadCreatePost)
	mux.HandleFunc("GET /thread/view/{id}", app.threadView)

	mux.HandleFunc("GET /thread/view/{id}/message/create", app.messageCreate)
	mux.HandleFunc("POST /thread/view/{id}/message/create", app.messageCreatePost)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}
