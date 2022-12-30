package handler

import (
	"net/http"
	"net/url"
	"text/template"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get a map that contains a key and value
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get id param value by Get func of url.Values
	idStr := values.Get("id")

	var tmpl *template.Template
	if idStr == "" || idStr == "1" {
		tmpl, err = template.ParseFiles("static/templates/home1.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	} else if idStr == "2" {
		tmpl, err = template.ParseFiles("static/templates/home2.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}
	tmpl.Execute(w, "")
}
