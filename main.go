package main

import (
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
