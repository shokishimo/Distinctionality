package handler

import (
	// "fmt"
	"log"
	"net/http"
	"text/template"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/img/home.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl.Execute(w, "")
	// fmt.Fprint(w, "Home page is requested")
}
