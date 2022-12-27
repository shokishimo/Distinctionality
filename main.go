package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/shokishimo/Distinctionality/db"
	server "github.com/shokishimo/Distinctionality/internal"
)

func main() {
	client, err := db.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Disconnect(client)
	srvMux := server.New()
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe("localhost:8080", srvMux); err != nil {
		log.Fatal("Error in running the server")
		fmt.Printf("e: %v\n", err)
	}
}
