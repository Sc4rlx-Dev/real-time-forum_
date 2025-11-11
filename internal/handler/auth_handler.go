package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/models"
	"real_time_forum/internal/repository"
	"strconv"
)

type Auth_handler struct {
	DB *sql.DB
}

func (h *Auth_handler) Register(w http.ResponseWriter, r *http.Request) {

    if err := r.ParseForm(); err != nil { http.Error(w,  "Failed to parse form" , http.StatusBadRequest) ; return}

    age, _ := strconv.Atoi(r.FormValue("Age"))

    user_data := models.User_data{
		Username:   r.FormValue("Username"),
		First_name: r.FormValue("First_name"),
		Last_name:  r.FormValue("Last_name"),
		Age:        age,
		Email:      r.FormValue("Email"),
		Password:   r.FormValue("Password"),
		Gender:     r.FormValue("Gender"),
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
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	login_data := models.Data{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
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

	// Set the secure session token
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session_token,
		Path:     "/",
		HttpOnly: false, 
	})

	// --- ADD THIS COOKIE ---
	// Set the username cookie for the frontend JavaScript to read
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    login_data.Username,
		Path:     "/",
		HttpOnly: false, // JavaScript CAN see this
	})
	// ---------------------

	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully"})
}