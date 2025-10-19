package repository

import (
	"database/sql"
	"log"
	"real_time_forum/internal/models"
	// "strconv"
)

func Get_user_id_by_username(db *sql.DB, username string) (int, error) {
	var user_id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&user_id)
	if err != nil {
		log.Printf("Error getting user ID for %s: %v", username, err)
		return 0, err
	}
	return user_id, nil
}

func Find_or_create_conversation(db *sql.DB, sender_id, receiver_id int) (int, error) {
	var conv_id int
	err := db.QueryRow(`
		SELECT id FROM conversations
		WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
		LIMIT 1`,
		sender_id, receiver_id, receiver_id, sender_id).Scan(&conv_id)

	if err == sql.ErrNoRows {
		res, err := db.Exec(`
			INSERT INTO conversations (sender_id, receiver_id) VALUES (?, ?)`,
			sender_id, receiver_id)
		if err != nil {
			log.Printf("Error creating conversation: %v", err)
			return 0, err
		}
		new_id, _ := res.LastInsertId()
		return int(new_id), nil
	} else if err != nil {
		log.Printf("Error querying conversation: %v", err)
		return 0, err
	}

	return conv_id, nil
}

func Insert_chat_message(db *sql.DB, msg models.Message, sender_id, receiver_id int) error {
	conv_id, err := Find_or_create_conversation(db, sender_id, receiver_id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO messages (message, conversation_id, sender_id, receiver_id, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		msg.Message, conv_id, sender_id, receiver_id, msg.Date)

	if err != nil {
		log.Printf("Error inserting message: %v", err)
	}
	return err
}
