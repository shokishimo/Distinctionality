package server

import (
	"net/http"

	handler "github.com/shokishimo/Distinctionality/internal/handler"
	// middleware "github.com/shokishimo/Distinctionality/internal/middleware"
)

// New creates a new HTTP server.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeFunc)
	mux.HandleFunc("/insertData", handler.CreateData)
	mux.HandleFunc("/get20Quizzes", handler.Get20Quizzes)
	mux.HandleFunc("/quiz", handler.Quiz)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return mux
}
