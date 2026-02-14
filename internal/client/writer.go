package client

import "log"

func (c *Client)WritePump(){
	defer c.Conn.Close()

	for{
		msg,ok:=<-c.Send
		if !ok{
			return
		}
		err:=c.Conn.WriteMessage(1,msg)
		if err!=nil{
			log.Println("Write error",err)
			return
		}
	}
}