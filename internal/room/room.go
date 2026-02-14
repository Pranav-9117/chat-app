package room

import (
	"log"
	"chat-app/internal/client"
)

type Room struct{
	ID string
	Clients map[*client.Client]bool
	Broadcast chan []byte
	Join chan *client.Client
	Leave chan *client.Client
}

func NewRoom(id string)*Room{
	return &Room{
		ID:id,
		Clients: make(map[*client.Client]bool),
		Broadcast: make(chan []byte),
		Join: make(chan *client.Client),
		Leave: make(chan *client.Client),
	}
}
func(r *Room)Run(){
	log.Println("Room started",r.ID)
	for{
		select{
		case cl := <-r.Join:
    		if !r.Clients[cl] {
        		r.Clients[cl] = true
        		log.Println("Client joined the room:", r.ID)
    		}
		
		case cl := <-r.Leave:
    		if r.Clients[cl] {
        		delete(r.Clients, cl)
        		close(cl.Send)
       		 log.Println("Client left the room:", r.ID)
    		}
		
		case msg:=<-r.Broadcast:
			for cl:=range r.Clients{
			select{
				case cl.Send<-msg:
				default:
					delete(r.Clients,cl)
					close(cl.Send)
				}

			}

		}
	}

}

func (r *Room) JoinClient(c *client.Client) {
	r.Join <- c
}

func (r *Room) BroadcastMessage(msg []byte) {
	r.Broadcast <- msg
}

func (r *Room) LeaveClient(c *client.Client) {
    r.Leave <- c
}
