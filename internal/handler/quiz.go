package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func Quiz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Quiz is accessed")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("static/templates/quiz.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, "")
}
