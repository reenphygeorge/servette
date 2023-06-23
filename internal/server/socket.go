package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/reenphygeorge/servette/internal/logger"
)

/*
	Global websocket connection variable.
	Holds connection data, if socket connection was successful.
	Also used in logger module.
*/
var GlobalConn *websocket.Conn

// Upgrades HTTP connection to a WebSocket connection.
func SocketUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	return conn, err
}

// To handle messages from client side.
func HandleMessage(conn *websocket.Conn) {
	GlobalConn = conn
	conn.WriteMessage(websocket.TextMessage, []byte("Welcome to Servette!"))

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

/*
	Requests the client to reload.
	Called on file change. 
*/
func ReloadRequest() {
	if GlobalConn != nil {
		serverMessage := []byte("Reload")
		err := GlobalConn.WriteMessage(websocket.TextMessage, serverMessage)
		if err != nil {
			logger.Error("")
		}
	}
}
