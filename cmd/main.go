package main

import ("log"
		"net/http"
		"chat-app/internal/server"
)

func main(){

	port:=":8080"

	srv:=server.NewServer()
	http.HandleFunc("/ws",srv.HandleWebSocket)
	log.Println("Server running on ",port)
	err:=http.ListenAndServe(port,nil)
	if err!=nil{
		log.Fatal("Server Failed",err)
	}

}