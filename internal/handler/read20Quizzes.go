package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/shokishimo/Distinctionality/model"
)

func Get20Quizzes(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var qandas []model.QandA
	qandas, err := model.Get20Quizzes()
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(qandas)

	fmt.Println("Get20 handle is done")
	return
}
