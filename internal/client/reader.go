package client

import (
	"chat-app/models"
	"encoding/json"
	"log"
)

func (c *Client) ReadPump(messageHandler func(models.Message)) {

	defer func() {
		log.Println("Client disconnected")
		c.Close()
	}()

	for {
		_, msgBytes, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var message models.Message
		err = json.Unmarshal(msgBytes, &message)
		if err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}

		messageHandler(message)
	}
}
