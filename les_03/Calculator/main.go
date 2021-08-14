package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Exps struct {
	expression string
}

func cal(w http.ResponseWriter, r *http.Request) {
	var exps Exps
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&exps)
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	fmt.Println(exps.expression + "asdas") //Khong ra Data cua Res
	result := Exps{expression: "123"}
	encoder.Encode(result)

}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandle).Methods("GET")
	r.HandleFunc("/cal", cal).Methods("POST")
	fs := http.FileServer(http.Dir("./templates/src"))
	r.PathPrefix("/src/").Handler(http.StripPrefix("/src", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
