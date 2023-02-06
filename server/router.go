package server

import (
	"wjjmjh/webrtc-signalling-golang/server/utils"

	"github.com/gorilla/mux"
)

func RouterWebsocketServices(fn utils.JWTAuthFunc, secret []byte) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//route.HandlerFunc, route.Name
	// Add routes needing Authentication
	for _, route := range routesProtectedWS {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(fn(route.HandlerFunc, secret))
	}

	// Add routes NOT needing Authentication
	for _, route := range routesOpenWS {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
