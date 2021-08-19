package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Exc(exps string) float64 {
	data := strings.Split(exps, "")
	result, _ := strconv.ParseFloat(data[0], 64)
	var resultA, resultB string
	var str string
	for i := range data {
		check := true
		if check {
			if data[i] != "+" || data[i] != "-" || data[i] != "*" || data[i] != "/" {
				resultA += data[i]
			} else {
				switch data[i] {
				case "+":
					str = "+"
					check = false
				case "-":
					str = "-"
					check = false
				case "*":
					str = "*"
					check = false
				case "/":
					str = "/"
					check = false
				}
			}

		} else {
			resultB += data[i]
		}
	}
	rs1, _ := strconv.ParseFloat(resultA, 8)
	rs2, _ := strconv.ParseFloat(resultB, 8)
	switch str {
	case "+":
		result = rs1 + rs2
	case "-":
		result = rs1 - rs2
	case "*":
		result = rs1 * rs2
	case "/":
		if rs2 == 0 {
			result = 0
		} else {
			result = rs1 / rs2
		}

	}
	return result
}

func cal(w http.ResponseWriter, r *http.Request) {
	var exps string
	err := json.NewDecoder(r.Body).Decode(&exps)
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}
	fmt.Println(exps)
	result := Exc(exps)
	json.NewEncoder(w).Encode(result)

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
