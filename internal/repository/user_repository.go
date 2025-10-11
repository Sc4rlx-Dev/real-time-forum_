package repository

import (
	"database/sql"
	"errors"
	// "hash"
	"real_time_forum/internal/models"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)


func Insert_user(db *sql.DB , usr *models.UserData) error{
	hashed_passwd ,err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
    if err != nil { return  errors.New("faild to hash passwd")}
	_, err = db.Exec(` INSERT INTO users (username, age, gender, firstname, lastname, password, email)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
    usr.Username, usr.Age, usr.Gender, usr.FirstName, usr.LastName, string(hashed_passwd), usr.Email)

    if err != nil { return errors.New("failed to register user: " + err.Error()) }
    return nil
}

func Auth_user(db *sql.DB, logindata *models.Data) (int, error){
	var store_pass string
	var user_id int

	err := db.QueryRow("SELECT id , password FROM users WHERE username = ? OR email = ?" ,
	logindata.Username , logindata.Password).Scan(&user_id , &store_pass)
	if err != nil { return 0, errors.New("invalid password")}
	return user_id , nil
}

func Create_session(db *sql.DB , user_id int , username string) (string, error){
	session_token , err := uuid.NewV4()
	if err != nil {
		return "" , err // i will check it later
	}
	_, err = db.Exec(`
		INSERT INTO sessions (user_id, username, session_id, expiry_date)
        VALUES (?, ?, ?, datetime('now', '+24 hours'))`,
		user_id, username, session_token.String())

		if err != nil { return "" , err }
return session_token.String() , nil
}