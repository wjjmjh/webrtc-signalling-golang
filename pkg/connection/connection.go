package connection

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// BufferSizes default:
// readBufferSize: 1024
// writeBufferSize: 1024
type BufferSizes struct {
	readBufferSize  int
	writeBufferSize int
}

// Creates buffers with default (1024/1024) read write sizes
func createDefaultBuffer() *BufferSizes {
	return &BufferSizes{1024, 1024}
}

// Creates buffers with custom read/write sizes
func createCustomBuffer(readSize int, writeSize int) *BufferSizes {
	return &BufferSizes{readSize, writeSize}
}

//makeUpgrader: make a websocket upgrader based on specified buffersizes.
// Utility function for UpgradeHTTPToWS function.
func makeUpgrader(bufferSizes *BufferSizes) websocket.Upgrader {

	if bufferSizes.readBufferSize == 0 {
		bufferSizes.readBufferSize = 1024
	}
	if bufferSizes.writeBufferSize == 0 {
		bufferSizes.writeBufferSize = 1024
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  bufferSizes.readBufferSize,
		WriteBufferSize: bufferSizes.writeBufferSize,
		CheckOrigin:     func(*http.Request) bool { return true },
	}
	return upgrader
}

// UpgradeHTTPToWS upgrades the HTTP server connection to the WebSocket protocol.
func UpgradeHTTPToWS(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := makeUpgrader(createDefaultBuffer())
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}
