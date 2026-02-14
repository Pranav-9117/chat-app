package server

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-app/internal/client"
	"chat-app/internal/room"
	"chat-app/models"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	manager  *room.Manager
}

func NewServer() *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		manager: room.NewManager(),
	}
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	log.Println("Connection established")

	cl := client.NewClient(conn)

	
	go cl.WritePump()

	go cl.ReadPump(func(msg models.Message) {

		
		rm := s.manager.GetRoom(msg.RoomId)

		
		rm.JoinClient(cl)

		log.Printf("Received message from %s in room %s: %s\n",
			msg.SenderId,
			msg.RoomId,
			msg.Msg,
		)

		
		finalMsg, _ := json.Marshal(msg)

		rm.BroadcastMessage(finalMsg)
	})
}
