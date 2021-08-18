package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Exps struct {
	Expression string `json:"expression"`
}

func cal(w http.ResponseWriter, r *http.Request) {
	exp := Exps{}
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		fmt.Fprint(w, "Error")
		return
	}
	fmt.Println(exp.Expression) //Khong ra Data cua Res
	result := Exps{Expression: exp.Expression}
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
