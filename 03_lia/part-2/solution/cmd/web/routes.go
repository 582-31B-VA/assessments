package main

import "net/http"

// routes returns a mux with all the registered routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /account/create", app.accountCreate)
	mux.HandleFunc("POST /account/create", app.accountCreatePOST)
	mux.HandleFunc("GET /account/view/{id}", app.accountView)

	mux.HandleFunc("GET /thread/create", app.threadCreate)
	mux.HandleFunc("POST /thread/create", app.threadCreatePOST)
	mux.HandleFunc("GET /thread/view/{id}", app.threadView)

	mux.HandleFunc("GET /thread/view/{id}/post/create", app.postCreate)
	mux.HandleFunc("POST /thread/view/{id}/post/create", app.postCreatePOST)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(commonHeaders(mux))
}
