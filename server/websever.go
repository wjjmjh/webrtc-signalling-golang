package server

import (
	"net/http"
	"wjjmjh/webrtc-signalling-golang/server/utils"

	"github.com/sirupsen/logrus"
)

// StartWebSocketServer StartWebServer starts a web server at the designated port.
func StartWebSocketServer(port string, secret []byte) {
	funcHandleAuth := utils.JWTAuthMiddleware
	r := RouterWebsocketServices(funcHandleAuth, secret)
	logrus.Infof("Starting WS service at %v", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logrus.Errorln("An error occurred starting HTTP listener at port " + port)
		logrus.Errorln("Error: " + err.Error())
	}
}
