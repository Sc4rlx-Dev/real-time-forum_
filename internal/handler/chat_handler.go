package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/repository"
	"strconv"
)

type Chat_handler struct {
	DB *sql.DB
}

// Get_users gets all users for the chat list
func (h *Chat_handler) Get_users(w http.ResponseWriter, r *http.Request) {
    // Get current user ID from middleware
    user_id, _ := strconv.Atoi(r.Header.Get("X-User-ID"))

    users, err := repository.Get_all_users(h.DB, user_id)
    if err != nil {
        http.Error(w, "Failed to get users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Get_messages gets chat history with a specific user
func (h *Chat_handler) Get_messages(w http.ResponseWriter, r *http.Request) {
    user_id1, _ := strconv.Atoi(r.Header.Get("X-User-ID"))
    
    // Get the username from the URL path, e.g., /api/messages/otheruser
    other_username := r.URL.Path[len("/api/messages/"):]
    if other_username == "" {
        http.Error(w, "Missing username in URL", http.StatusBadRequest)
        return
    }

    user_id2, err := repository.Get_user_id_by_username(h.DB, other_username)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    messages, err := repository.Get_messages(h.DB, user_id1, user_id2)
    if err != nil {
        http.Error(w, "Failed to get messages", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}