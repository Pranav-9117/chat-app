package main

import (
	"chat-app/internal/server"
	"log"
	"net/http"
)

func main() {

	port := ":8080"

	srv := server.NewServer()
	http.HandleFunc("/ws", srv.HandleWebSocket)
	log.Println("Server running on ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server Failed", err)
	}

}
