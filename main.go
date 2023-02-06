package main

import (
	"strconv"
	"wjjmjh/webrtc-signalling-golang/models"
	"wjjmjh/webrtc-signalling-golang/server"
	"wjjmjh/webrtc-signalling-golang/server/handlers"

	"github.com/spf13/viper"
)

const appPort int = 4000

func init() {
	// initialise services
	viper.SetDefault("websocket_port", strconv.Itoa(appPort))
}

func main() {
	// initialise websocket clients pool
	pool := models.NewPool()
	// start websocket clients pool to receive messages
	go pool.Start()
	handlers.Pool = pool

	// TODO: implement logics that obtain the secret here
	var secret []byte
	server.StartWebSocketServer(viper.GetString("websocket_port"), secret)
}
