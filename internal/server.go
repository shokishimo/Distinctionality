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
	mux.HandleFunc("/home", handler.HomeFunc)
	mux.HandleFunc("/insert", handler.CreateData)
	mux.HandleFunc("/get20Quizzes", handler.Get20Quizzes)
	// mux.HandleFunc("/login", handler.Login)
	// mux.HandleFunc("/logout", handler.Logout)
	// mux.HandleFunc("/signup", handler.Signup)
	// mux.HandleFunc("/profile", handler.Profile)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // or http.Dir("static")
	return mux
}
