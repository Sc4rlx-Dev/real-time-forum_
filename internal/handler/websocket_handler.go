package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"real_time_forum/internal/models"
	"real_time_forum/internal/repository"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[string]*websocket.Conn)
var clients_mutex = sync.Mutex{}

func Ws_handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		username := r.URL.Query().Get("username")
		if username == "" {
			log.Println("WebSocket connection rejected: missing username")
			return
		}

		clients_mutex.Lock()
		clients[username] = conn
		clients_mutex.Unlock()
		log.Printf("WebSocket client connected: %s", username)

		defer func() {
			clients_mutex.Lock()
			delete(clients, username)
			clients_mutex.Unlock()
			log.Printf("WebSocket client disconnected: %s", username)
		}()

		for {
			_, message_bytes, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error from %s: %v", username, err)
				break
			}

			var msg models.Message
			if err := json.Unmarshal(message_bytes, &msg); err != nil {
				log.Printf("JSON unmarshal error from %s: %v", username, err)
				continue
			}

			if msg.From != username {
				log.Printf("Message 'From' field mismatch for user %s", username)
				continue
			}

			sender_id, err_s := repository.Get_user_id_by_username(db, msg.From)
			receiver_id, err_r := repository.Get_user_id_by_username(db, msg.To)

			if err_s != nil || err_r != nil {
				log.Printf("Could not find sender or receiver ID for message from %s to %s", msg.From, msg.To)
				continue
			}

			err = repository.Insert_chat_message(db, msg, sender_id, receiver_id)
			if err != nil {
				log.Printf("Failed to save message from %s: %v", username, err)
				continue
			}

			clients_mutex.Lock()
			recipient_conn, is_online := clients[msg.To]
			clients_mutex.Unlock()

			if is_online {
				if err := recipient_conn.WriteJSON(msg); err != nil {
					log.Printf("Error sending message to %s: %v", msg.To, err)
				}
			} else {
				log.Printf("Recipient %s is offline. Message saved.", msg.To)
			}
		}
	}
}
