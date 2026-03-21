package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == "GET" {
		http.Redirect(w, r, "/form.html", http.StatusMovedPermanently)
		return
	}

	fmt.Fprintf(w, "Post req was successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", address)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotAcceptable)
		return
	}

	fmt.Fprintf(w, "Hello and welcome to this website!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	// http.FileServer: built-in static file server
	// Browser → Go server → returns files from your disk
	// http.Dir: Converts a directory into a FileSystem interface
	// returns a handler that serves HTTP requests

	// http.Handle     → takes a Handler
	// http.HandleFunc → takes a function. Go internally converts your function into a handler.

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", sayHello)

	PORT := 8080
	fmt.Printf("Start serving on port: %v/n", PORT)

	// nil tells Go - Use the default request router
	// If you pass:
	// http.ListenAndServe(":8080", myHandler)
	// Then http.HandleFunc routes are ignored. Else for all pages there will be only one handler
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
