package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandle).Methods("GET")
	fs := http.FileServer(http.Dir("./templates/src"))
	r.PathPrefix("/src/").Handler(http.StripPrefix("/src", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
