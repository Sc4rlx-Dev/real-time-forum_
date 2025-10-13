package repository

import (
	"database/sql"
	"errors"
	"real_time_forum/internal/models"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Corrected: Capitalized to make it public
func Insert_user(db *sql.DB, usr *models.User_data) error {
	hashed_passwd, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
    if err != nil { 
        return errors.New("failed to hash password")
    }

	_, err = db.Exec(`INSERT INTO users (username, age, gender, firstname, lastname, password, email)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
        usr.Username, usr.Age, usr.Gender, usr.First_name, usr.Last_name, string(hashed_passwd), usr.Email)

    if err != nil { 
        return errors.New("failed to register user: " + err.Error()) 
    }
    return nil
}

// Corrected: Capitalized to make it public
func Auth_user(db *sql.DB, login_data *models.Data) (int, error) {
	var stored_pass string
	var user_id int

	err := db.QueryRow("SELECT id, password FROM users WHERE username = ? OR email = ?",
		login_data.Username, login_data.Username).Scan(&user_id, &stored_pass)
	if err != nil {
		return 0, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(stored_pass), []byte(login_data.Password))
	if err != nil {
		return 0, errors.New("invalid password")
	}

	return user_id, nil
}

// Corrected: Capitalized to make it public
func Create_session(db *sql.DB, user_id int, username string) (string, error) {
	session_token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
    _, err = db.Exec("DELETE FROM sessions WHERE user_id = ?", user_id)
    if err != nil {
        return "", err
    }

	_, err = db.Exec(`INSERT INTO sessions (user_id, username, session_id, expiry_date)
        VALUES (?, ?, ?, datetime('now', '+24 hours'))`,
		user_id, username, session_token.String())

	if err != nil { 
        return "", err 
    }
	return session_token.String(), nil
}