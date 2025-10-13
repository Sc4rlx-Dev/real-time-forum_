package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "real_time_forum/internal/models"
    "real_time_forum/internal/repository"
)

type AuthHandler struct {
    DB *sql.DB
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

    var userData = models.UserData{
        Username:  "test",
        FirstName: "Test",
        LastName:  "User",
        Age:       23,
        Email:     "test@example.com",
        Password:  "password123",
        Gender:    "Male",
    }

    err := repository.Insert_user(h.DB, &userData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var loginData models.Data
    loginData = models.Data{
        Username: "testuser",
        Password: "password123",
    }
    userID, err := repository.Auth_user(h.DB, &loginData)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    sessionToken, err := repository.Create_session(h.DB, userID, loginData.Username)
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "session_token",
        Value:    sessionToken,
        HttpOnly: true,
    })
    json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully"})
}