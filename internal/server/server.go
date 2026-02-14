package server

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"chat-app/internal/client"
	"chat-app/internal/room"
)

type Server struct {
	upgrader websocket.Upgrader
	room *room.Room
}

func NewServer() *Server {
	r:=room.NewRoom("Valorant")
	go r.Run()
	return &Server{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		room: r,
	}
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed", err)
		return
	}
	log.Println("Connection established")
	cl := client.NewClient(conn)
	s.room.Join<-cl
	go cl.ReadPump(s.room.Broadcast)
	go cl.WritePump()
}
