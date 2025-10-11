package repository

import (
	"database/sql"
	"errors"
	"hash"
	"real_time_forum/internal/models" // Import your models

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