package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func SocketUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn,error){
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	return conn, err
}

func HandleMessage(conn *websocket.Conn) {
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket read error:", err)
            return
        }
        log.Printf("Received message: %s", message)

        // Echo the message back to the client
		serverMessage := []byte("Reload")
        err = conn.WriteMessage(messageType, serverMessage)
        if err != nil {
            log.Println("WebSocket write error:", err)
            return
        }
    }
}
