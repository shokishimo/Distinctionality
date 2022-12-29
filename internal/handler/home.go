package handler

import (
	"net/http"
	"net/url"
	"strconv"
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

	// simple "/" with no param -> return home page
	if idStr == "" {
		tmpl, err := template.ParseFiles("static/templates/home1.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		tmpl.Execute(w, "")
	}

	// Otherwise get id param value of type String and convert it into int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var tmpl *template.Template
	if id == 1 {
		tmpl, err = template.ParseFiles("static/templates/home1.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	} else if id == 2 {
		tmpl, err = template.ParseFiles("static/templates/home2.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}
	tmpl.Execute(w, "")
}
