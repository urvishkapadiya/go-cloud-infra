package main

import (
	"log"
	"net/http"
)

func main() {
	// File server for static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Handle dynamic routes if needed in the future
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/project", projectHandler)
	http.HandleFunc("/photography", photographyHandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/project.html")
}

func photographyHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/photography.html")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/blog.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}
