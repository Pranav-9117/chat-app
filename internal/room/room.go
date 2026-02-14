package room

import (
	"log"
	"chat-app/internal/client"
	"os"
	"path/filepath"
	"bufio"
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
func (r *Room) loadHistory(c *client.Client) {

	filePath := "storage/" + r.ID + ".txt"

	file, err := os.Open(filePath)
	if err != nil {
		
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		
		c.Send <- append([]byte{}, line...)
	}
}

func (r *Room) saveMessage(msg []byte) {
    os.MkdirAll("storage", os.ModePerm)

    filePath := filepath.Join("storage", r.ID+".txt")

    f, err := os.OpenFile(filePath,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("File write error:", err)
        return
    }
    defer f.Close()

    f.Write(msg)
    f.Write([]byte("\n"))
}

func(r *Room)Run(){
	log.Println("Room started",r.ID)
	for{
		select{
		case cl := <-r.Join:
    		if !r.Clients[cl] {
        		r.Clients[cl] = true
        		log.Println("Client joined the room:", r.ID)
				go r.loadHistory(cl)
    		}
		
		case cl := <-r.Leave:
    		if r.Clients[cl] {
        		delete(r.Clients, cl)
        		close(cl.Send)
       		 log.Println("Client left the room:", r.ID)
    		}
		
		case msg:=<-r.Broadcast:
			r.saveMessage(msg)
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
