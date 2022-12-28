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
	w.Header().Set("Content-Type", "application/json")

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

	// Print the unmarshalled slice of QandA struct
	fmt.Println(qas)

	err = model.CreateQandAs(qas)
	if err != nil {
		fmt.Println("wasn't able to insert")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Done with inserting new data!")
	return
}
