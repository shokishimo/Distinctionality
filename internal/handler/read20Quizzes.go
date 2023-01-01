package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	model "github.com/shokishimo/Distinctionality/model"
)

func Get20Quizzes(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is GET
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
	version := values.Get("version")
	level := values.Get("level")

	var qandas []model.QandA
	qandas, err = model.Get20Quizzes(version, level)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qandas)

	return
}
