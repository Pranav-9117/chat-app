package client

import "log"

func (c *Client)ReadPump(){
	defer func(){
		log.Println("Client disconnected")
		c.Conn.Close()
	}()
	for{
		_,msg,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println("Read error",err)
			break
		}
		log.Printf("Received message: %s",msg)

	}
	
}