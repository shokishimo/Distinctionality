package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/shokishimo/Distinctionality/internal"
)

func main() {
	srvMux := server.New()
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe("localhost:8080", srvMux); err != nil {
		log.Fatal("Error in running the server")
		fmt.Printf("e: %v\n", err)
	}
}
