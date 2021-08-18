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
	var exp Exps
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&exp)
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}
	
	fmt.Println(exp.expression ) //Khong ra Data cua Res
	result := Exps{expression: exp.expression}
	encoder := json.NewEncoder(w)
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
