package client

import(
	"log"
	"github.com/gorilla/websocket"
)

type Client struct{
	ID string
	Conn *websocket.Conn
	Send chan []byte
	Room interface{}
}
func NewClient(conn *websocket.Conn)*Client{
	return &Client{
		Conn:conn,
		Send: make(chan []byte,256),


	}
}
func (c *Client)Close(){
	err:=c.Conn.Close()
	if err!=nil{
		log.Println("failed to close the connection")

	}
	close(c.Send)
}