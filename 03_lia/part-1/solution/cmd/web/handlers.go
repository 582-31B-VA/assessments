package main

import (
	"net/http"
)

// home shows the 10 latest threads.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("home", w)
}

// accountCreate shows a form the create an account.
func (app *application) accountCreate(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("account-create", w)
}

// accountCreatePost creates an account with the info in the POST request.
func (app *application) accountCreatePost(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
}

// accountView shows information about an account.
func (app *application) accountView(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("account-view", w)
}

// threadCreate shows a form to create a thread.
func (app *application) threadCreate(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("thread-create", w)
}

// threadCreatePost creates a post with the info in the POST request.
func (app *application) threadCreatePost(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
}

// threadView shows a thread.
func (app *application) threadView(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("thread-view", w)
}

// messageCreate shows a form to create a message.
func (app *application) messageCreate(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
	app.renderTemplate("message-create", w)
}

// messageCreatePost creates a message with the info in the POST request.
func (app *application) messageCreatePost(w http.ResponseWriter, r *http.Request) {
	app.logRequest(r)
}
