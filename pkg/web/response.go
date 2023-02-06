package web

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJSONResponse(w http.ResponseWriter, status int, msgType string, messages ...string) {
	resp := make(map[string][]string)
	resp[msgType] = messages
	retJSON, err := json.Marshal(resp)
	if err != nil {
		retJSON = []byte("{\"error\": \"error message creation failed\" }")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(retJSON)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	_, err = w.Write(retJSON)
	if err != nil {
		return
	}
}

func WriteInternalServerErrorJsonResponse(w http.ResponseWriter, messages ...string) {
	WriteJSONResponse(w, http.StatusInternalServerError,
		"error",
		messages...)
}
