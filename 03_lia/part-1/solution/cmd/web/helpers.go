package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// logRequest logs the request.
func (app *application) logRequest(r *http.Request) {
	app.logger.Info("request received", "method", r.Method, "uri", r.URL.Path)
}

// serverError writes a log entry at Error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and corresponding description
// to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// renderTemplate parses and renders the template with the given name to the
// given ResponseWriter.
func (app *application) renderTemplate(name string, w http.ResponseWriter) error {
	ts, err := template.ParseFiles(
		"./ui/html/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/nav.html",
		"./ui/html/partials/footer.html",
		"./ui/html/pages/"+name+".html",
	)
	if err != nil {
		return fmt.Errorf("parsing template %s: %w", name, err)
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		return fmt.Errorf("executing template %s: %w", name, err)
	}
	return nil
}
