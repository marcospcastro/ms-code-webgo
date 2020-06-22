package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Greeting(t string) (s string) {
	s = "<b>"+t+"</b>"
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	msg := Greeting("Code.education Rocks!")
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "index.html", template.HTML(msg))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
}