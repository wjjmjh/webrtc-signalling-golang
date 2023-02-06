package server

import (
	"net/http"
	"wjjmjh/webrtc-signalling-golang/server/handlers"
	"wjjmjh/webrtc-signalling-golang/server/utils"
)

// Routes defines the type Routes which is just an array (slice) of model.Route structs.
type Routes []utils.Route

var routesProtectedWS = Routes{}

var routesOpenWS = Routes{
	utils.Route{
		Name:        "WebSocket",
		Method:      "GET",
		Pattern:     "/connect",
		HandlerFunc: handlers.AcceptConnection,
	},
	utils.Route{
		Name:    "HealthCheck",
		Method:  "GET",
		Pattern: "/healthz",
		HandlerFunc: func(writer http.ResponseWriter, r *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, err := writer.Write([]byte("OK"))
			if err != nil {
				return
			}
		},
	},

	utils.Route{
		Name:    "ReadinessCheck",
		Method:  "GET",
		Pattern: "/readyz",
		HandlerFunc: func(writer http.ResponseWriter, r *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, err := writer.Write([]byte("OK"))
			if err != nil {
				return
			}
		},
	},
}
