package main

import (
	"errors"
	"forum/internal/models"
	"net/http"
	"strconv"
)

// home shows the 10 latest threads.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	threads, err := app.threads.Latests()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := app.newTemplateData()
	data.Threads = threads
	app.render(w, r, http.StatusOK, "home", data)
}

// accountCreate shows a form the create an account.
func (app *application) accountCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "account-create", app.newTemplateData())
}

// accountCreatePOST creates an account with the info in the POST request.
func (app *application) accountCreatePOST(w http.ResponseWriter, r *http.Request) {
}

// accountView shows information about an account.
func (app *application) accountView(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "account-view", app.newTemplateData())
}

// threadCreate shows a form to create a thread.
func (app *application) threadCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "thread-create", app.newTemplateData())
}

// threadCreatePOST creates a post with the info in the POST request.
func (app *application) threadCreatePOST(w http.ResponseWriter, r *http.Request) {
	_, err := app.threads.Insert("Test", 1)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// threadView shows a thread.
func (app *application) threadView(w http.ResponseWriter, r *http.Request) {
	idSegment := r.PathValue("id")
	id, err := strconv.Atoi(idSegment)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	thread, err := app.threads.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData()
	data.Thread = thread

	app.render(w, r, http.StatusOK, "thread-view", data)
}

// postCreate shows a form to create a message.
func (app *application) postCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "message-create", app.newTemplateData())
}

// postCreatePOST creates a message with the info in the POST request.
func (app *application) postCreatePOST(w http.ResponseWriter, r *http.Request) {
	idSegment := r.PathValue("id")
	id, err := strconv.Atoi(idSegment)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	ThreadId, err := strconv.Atoi(idSegment)
	if err != nil || ThreadId < 1 {
		http.NotFound(w, r)
		return
	}
	app.posts.Insert("test", ThreadId, 1)
}
