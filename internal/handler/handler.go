package handler

import (
	"fmt"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page is requested")
}
