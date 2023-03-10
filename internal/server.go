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
	mux.HandleFunc("/endQuiz", handler.EndQuiz)
	// http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "favicon.ico")
	// })

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/404", http.NotFound)
	return mux
}
