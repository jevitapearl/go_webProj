package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()                               // initialize a servemux
	fileServer := http.FileServer(http.Dir("./ui/static/")) // Serves all files in ./ui/static

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) // anything starting with /static/ will be handled

	mux.HandleFunc("/", app.home) // register home() for "/" path
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/snippet/view", app.viewSnippet)

	return mux
}
