package models

// -------- message actions ----------

const Broadcaster = "broadcaster"
const Connect = "connect"
const Watcher = "watcher"
const Offer = "offer"
const Answer = "answer"
const Candidate = "candidate"
const Disconnect = "disconnect"
const DisconnectPeer = "disconnectPeer"

type Message struct {
	From           string      `json:"from"`
	To             string      `json:"to"`
	MessageType    string      `json:"messageType"`
	MessageContent interface{} `json:"messageContent"`
}
