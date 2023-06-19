package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/reenphygeorge/light-server/internal/logger"
)

var globalConn *websocket.Conn

func SocketUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn,error){
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	return conn, err
}

func HandleMessage(conn *websocket.Conn) {
    globalConn = conn
    conn.WriteMessage(websocket.TextMessage, []byte("Welcome to Light Server"))

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func ReloadRequest() {
	if globalConn != nil {
		serverMessage := []byte("Reload")
		err := globalConn.WriteMessage(websocket.TextMessage, serverMessage)
		if err != nil {
			logger.Error()
		}
	}
}
