package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/models"
	"real_time_forum/internal/repository"
	"regexp"
	"strconv"
	"strings"
)

type Auth_handler struct {
	DB *sql.DB
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func validate_registration(user *models.User_data) []ValidationError {
	var errors []ValidationError

	// Username validation
	if len(user.Username) < 3 || len(user.Username) > 20 {
		errors = append(errors, ValidationError{
			Field:   "Username",
			Message: "Username must be between 3 and 20 characters",
		})
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(user.Username) {
		errors = append(errors, ValidationError{
			Field:   "Username",
			Message: "Username can only contain letters, numbers, and underscores",
		})
	}

	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		errors = append(errors, ValidationError{
			Field:   "Email",
			Message: "Invalid email format",
		})
	}

	// Password validation
	if len(user.Password) < 6 {
		errors = append(errors, ValidationError{
			Field:   "Password",
			Message: "Password must be at least 6 characters long",
		})
	}

	// Name validation
	if len(user.First_name) < 2 || len(user.First_name) > 50 {
		errors = append(errors, ValidationError{
			Field:   "First_name",
			Message: "First name must be between 2 and 50 characters",
		})
	}
	if len(user.Last_name) < 2 || len(user.Last_name) > 50 {
		errors = append(errors, ValidationError{
			Field:   "Last_name",
			Message: "Last name must be between 2 and 50 characters",
		})
	}

	// Age validation
	if user.Age < 13 || user.Age > 120 {
		errors = append(errors, ValidationError{
			Field:   "Age",
			Message: "Age must be between 13 and 120",
		})
	}

	// Gender validation
	validGenders := []string{"male", "female", "other", "prefer not to say"}
	genderValid := false
	for _, g := range validGenders {
		if strings.ToLower(user.Gender) == g {
			genderValid = true
			break
		}
	}
	if !genderValid {
		errors = append(errors, ValidationError{
			Field:   "Gender",
			Message: "Invalid gender selection",
		})
	}

	return errors
}

func (h *Auth_handler) Register(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(r.FormValue("Age"))
	if err != nil {
		age = 0
	}

	user_data := models.User_data{
		Username:   r.FormValue("Username"),
		First_name: r.FormValue("First_name"),
		Last_name:  r.FormValue("Last_name"),
		Age:        age,
		Email:      r.FormValue("Email"),
		Password:   r.FormValue("Password"),
		Gender:     r.FormValue("Gender"),
	}

	// Validate input
	validationErrors := validate_registration(&user_data)
	if len(validationErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ValidationErrors{Errors: validationErrors})
		return
	}

	err = repository.Insert_user(h.DB, &user_data)
	if err != nil {
		// Check if it's a unique constraint error
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Username or email already exists",
			})
			return
		}
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
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

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session_token,
		Path:     "/",
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    login_data.Username,
		Path:     "/",
		HttpOnly: false,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Logged in successfully",
		"username": login_data.Username,
	})
}

func (h *Auth_handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == nil && cookie.Value != "" {
		// Delete session from database
		repository.Delete_session(h.DB, cookie.Value)
	}

	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}
