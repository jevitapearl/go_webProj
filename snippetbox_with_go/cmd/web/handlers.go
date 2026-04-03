package main

import (
	"fmt"
	"net/http"
	"strconv"

	"html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...) // read template files into a set
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil) // Implementing the template set to write content

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/create" {
		app.notFound(w)
		return
	}

	if r.Method != "POST" {
		// w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// fmt.Fprint(w, "Method not allowed\n")

		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Creating a new snippet...\n")
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with id %v", id)
}
