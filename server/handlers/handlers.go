package handlers

import (
	"net/http"
	"wjjmjh/webrtc-signalling-golang/models"
	"wjjmjh/webrtc-signalling-golang/pkg/connection"
	"wjjmjh/webrtc-signalling-golang/pkg/web"
)

var Pool *models.Pool

// AcceptConnection Websocket client handler. Blocks on ReadMessage
// Any message will cause client to be disconnected
// Attempts to close gracefully by sending client close message and websocket.CloseNormalClosure (1000) with text "closing"
// Waits for a grace period, then closes connection and removes client from Map of connections
func AcceptConnection(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.URL.Query()["userId"]

	if !ok || len(userId[0]) < 1 {
		web.WriteInternalServerErrorJsonResponse(w, "Url Param 'userId' is missing")
		return
	}

	wsConnection, err := connection.UpgradeHTTPToWS(w, r)
	if err != nil {
		web.WriteInternalServerErrorJsonResponse(
			w,
			"Error in establishing websocket connection: ",
			err.Error(),
		)
	}

	// create a new client
	client := models.CreateClient(userId[0], wsConnection, Pool)

	// make client listening for new messages
	go client.Read()

	// register
	Pool.Register <- client

}
