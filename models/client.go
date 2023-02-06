package models

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func CreateClient(userId string, conn *websocket.Conn, pool *Pool) *Client {
	logrus.Info("creating client", userId)
	return &Client{userId, conn, pool}
}

func (client *Client) Disconnect() {
	client.Pool.Unregister <- client
	err := client.Conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (client *Client) Read() {
	defer client.Disconnect()

	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		err = client.HandleNewMessage(p)
		if err != nil {
			return
		}
	}
}

// ---------- client logics ----------

func (client *Client) FindAndSend(to string, message interface{}) error {
	for c, _ := range client.Pool.Clients {
		if to == c.ID {
			if err := c.Conn.WriteJSON(message); err != nil {
				log.Println("Error with writing json message to client's connection", err)
				return err
			}
		}
	}
	return nil
}

func (client *Client) SendOthers(message interface{}) error {
	for c, _ := range client.Pool.Clients {
		if c.ID != client.ID {
			if err := c.Conn.WriteJSON(message); err != nil {
				log.Println("Error with writing json message to client's connection", err)
				return err
			}
		}
	}
	return nil
}

func (client *Client) HandleNewMessage(messageByte []byte) error {
	var unmarshalledMessage Message
	err := json.Unmarshal(messageByte, &unmarshalledMessage)
	if err != nil {
		log.Println("Error with unmarshal", err)
		return err
	}

	switch unmarshalledMessage.MessageType {
	case Broadcaster:
		logrus.Info(Broadcaster, unmarshalledMessage)
		// broadcast to all clients
		err := client.SendOthers(unmarshalledMessage)
		if err != nil {
			return err
		}
	case Connect:
		logrus.Info(Connect, unmarshalledMessage)
		// broadcast to all clients
		err := client.SendOthers(unmarshalledMessage)
		if err != nil {
			return err
		}
	case Watcher:
		logrus.Info(Watcher, unmarshalledMessage)
		// inform the broadcaster for a new watcher
		err := client.FindAndSend(unmarshalledMessage.To, unmarshalledMessage)
		if err != nil {
			return err
		}
	case Offer:
		logrus.Info(Offer, unmarshalledMessage)
		// broadcaster offers remote description
		err := client.SendOthers(unmarshalledMessage)
		if err != nil {
			return err
		}
	case Answer:
		logrus.Info(Answer, unmarshalledMessage)
		// inform the broadcaster that one watcher answers
		err := client.FindAndSend(unmarshalledMessage.To, unmarshalledMessage)
		if err != nil {
			return err
		}
	case Candidate:
		logrus.Info(Candidate, unmarshalledMessage)
		// inform the broadcaster for one new ICE candidate
		err := client.FindAndSend(unmarshalledMessage.To, unmarshalledMessage)
		if err != nil {
			return err
		}
	case DisconnectPeer:
		logrus.Info(DisconnectPeer, unmarshalledMessage)
		err := client.FindAndSend(unmarshalledMessage.To, unmarshalledMessage)
		if err != nil {
			return err
		}
	case Disconnect:
		logrus.Info(Disconnect, unmarshalledMessage)
		client.Disconnect()
	}

	return nil
}
