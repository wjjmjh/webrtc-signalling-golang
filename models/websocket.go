package models

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type PubSubData struct {
	Type      string `mapstructure:"type"`
	ID        string `mapstructure:"id"`
	Status    string `mapstructure:"status"`
	StatusMsg string `mapstructure:"status_message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
