package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"real_time_forum/internal/models"
	"real_time_forum/internal/repository"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	clients       = make(map[string]*websocket.Conn)
	clients_mutex = sync.Mutex{}
)

// Broadcast_user_list sends the current list of online users to everyone
func Broadcast_user_list() {
	clients_mutex.Lock()
	defer clients_mutex.Unlock()

	var online_users []string
	for user := range clients {
		online_users = append(online_users, user)
	}

	msg := models.Message{
		Type: "user_list",
		Data: online_users,
	}

	for user, conn := range clients {
		if err := conn.WriteJSON(msg); err != nil {
			log.Printf("Error broadcasting user list to %s: %v", user, err)
		}
	}
}

func Ws_handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println("WebSocket connection rejected: no session cookie")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user_id, username, err := repository.Get_user_from_session(db, cookie.Value)
		if err != nil {
			log.Println("WebSocket connection rejected: invalid session")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		clients_mutex.Lock()
		clients[username] = conn
		clients_mutex.Unlock()
		log.Printf("WebSocket client connected: %s", username)

		Broadcast_user_list()

		defer func() {
			clients_mutex.Lock()
			delete(clients, username)
			clients_mutex.Unlock()
			log.Printf("WebSocket client disconnected: %s", username)
			Broadcast_user_list()
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

			if msg.From_username != username {
				log.Printf("Message 'From' field mismatch for user %s. Expected %s, got %s", username, username, msg.From_username)
				continue
			}

			receiver_id, err_r := repository.Get_user_id_by_username(db, msg.To_username)
			if err_r != nil {
				log.Printf("Could not find receiver ID for %s", msg.To_username)
				continue
			}

			err = repository.Insert_chat_message(db, msg, user_id, receiver_id)
			if err != nil {
				log.Printf("Failed to save message from %s: %v", username, err)
				continue
			}

			clients_mutex.Lock()
			recipient_conn, is_online := clients[msg.To_username]
			clients_mutex.Unlock()

			if is_online {
				if err := recipient_conn.WriteJSON(msg); err != nil {
					log.Printf("Error sending message to %s: %v", msg.To_username, err)
				}
			} else {
				log.Printf("Recipient %s is offline. Message saved.", msg.To_username)
			}
		}
	}
}
