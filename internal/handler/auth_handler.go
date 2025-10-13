package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "real_time_forum/internal/models"
    "real_time_forum/internal/repository"
)

type Auth_handler struct {
	DB *sql.DB
}

func (h *Auth_handler) Register(w http.ResponseWriter, r *http.Request) {
    // Corrected: Merged declaration and assignment
	user_data := models.User_data{
		Username:   "test",
		First_name: "Test",
        Last_name:  "User",
        Age:        23,
        Email:      "test@example.com",
        Password:   "password123",
        Gender:     "Male",
	}
	err := repository.Insert_user(h.DB, &user_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func (h *Auth_handler) Login(w http.ResponseWriter, r *http.Request) {
    login_data := models.Data{
        Username: "test",
        Password: "password123",
    }

    user_id, err := repository.Auth_user(h.DB, &login_data)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    session_token, err := repository.Create_session(h.DB, user_id, login_data.Username)
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "session_token",
        Value:    session_token,
        HttpOnly: true,
    })
    json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully"})
}