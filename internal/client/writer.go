package client

import (
	"log"

	"github.com/gorilla/websocket"
)

func (c *Client)WritePump(){
	defer c.Conn.Close()

	for{
		msg,ok:=<-c.Send
		if !ok{
			return
		}
		err:=c.Conn.WriteMessage(websocket.TextMessage,msg)
		if err!=nil{
			log.Println("Write error",err)
			return
		}
	}
}