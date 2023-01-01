package handler

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	model "github.com/shokishimo/Distinctionality/model"
)

func CreateData(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get params in the header by Get func of url.Values
	versionStr := r.Header.Get("version")
	levelStr := r.Header.Get("level")
	if versionStr == "" || levelStr == "" {
		fmt.Println("Error: Values.Get")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	versionStr = "Distinctionality" + versionStr
	levelStr = "level" + levelStr

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the request body of json into a slice of QandA struct
	var qas []model.QandA
	err = json.Unmarshal(body, &qas)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	err = model.CreateQandAs(qas, versionStr, levelStr)
	if err != nil {
		w.Write([]byte("Failed to insert data"))
		return
	}

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Succeed: Inserted a given data\n")
	return
}
