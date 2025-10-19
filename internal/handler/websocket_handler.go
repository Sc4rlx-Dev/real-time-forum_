package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Ws_handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		log.Printf("Received message type %d: %s\n", messageType, p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
	log.Println("WebSocket client disconnected")
}
