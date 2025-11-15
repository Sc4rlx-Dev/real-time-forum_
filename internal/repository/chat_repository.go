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

func Get_messages(db *sql.DB, user_id1 int, user_id2 int) ([]models.Message, error) {
	rows, err := db.Query(`
        SELECT id, message, sender_id, receiver_id, created_at
        FROM messages
        WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
        ORDER BY created_at ASC`,
		user_id1, user_id2, user_id2, user_id1)

	if err != nil {
		log.Printf("Error getting messages: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		var sender_id, receiver_id int
		if err := rows.Scan(&msg.ID, &msg.Message, &sender_id, &receiver_id, &msg.Date); err != nil {
			log.Printf("Error scanning message: %v", err)
			continue
		}

		if sender_id == user_id1 {
			msg.From_username, _ = Get_username_by_id(db, user_id1)
			msg.To_username, _ = Get_username_by_id(db, user_id2)
		} else {
			msg.From_username, _ = Get_username_by_id(db, user_id2)
			msg.To_username, _ = Get_username_by_id(db, user_id1)
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func Get_messages_paginated(db *sql.DB, user_id1 int, user_id2 int, before string, limit int) ([]models.Message, error) {
	var rows *sql.Rows
	var err error

	if before != "" {
		// Get messages before a specific ID
		rows, err = db.Query(`
			SELECT id, message, sender_id, receiver_id, created_at
			FROM messages
			WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
			AND id < ?
			ORDER BY created_at DESC
			LIMIT ?`,
			user_id1, user_id2, user_id2, user_id1, before, limit)
	} else {
		// Get the latest messages
		rows, err = db.Query(`
			SELECT id, message, sender_id, receiver_id, created_at
			FROM messages
			WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
			ORDER BY created_at DESC
			LIMIT ?`,
			user_id1, user_id2, user_id2, user_id1, limit)
	}

	if err != nil {
		log.Printf("Error getting paginated messages: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		var sender_id, receiver_id int
		if err := rows.Scan(&msg.ID, &msg.Message, &sender_id, &receiver_id, &msg.Date); err != nil {
			log.Printf("Error scanning message: %v", err)
			continue
		}

		if sender_id == user_id1 {
			msg.From_username, _ = Get_username_by_id(db, user_id1)
			msg.To_username, _ = Get_username_by_id(db, user_id2)
		} else {
			msg.From_username, _ = Get_username_by_id(db, user_id2)
			msg.To_username, _ = Get_username_by_id(db, user_id1)
		}
		messages = append(messages, msg)
	}

	// Reverse to get chronological order
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func Get_conversations(db *sql.DB, user_id int) ([]models.Conversation, error) {
	// Get all users the current user has conversations with, along with their last message
	rows, err := db.Query(`
		SELECT 
			u.id,
			u.username,
			u.firstname,
			u.lastname,
			m.message as last_message,
			m.created_at as last_message_at
		FROM users u
		LEFT JOIN (
			SELECT 
				CASE 
					WHEN sender_id = ? THEN receiver_id
					ELSE sender_id
				END as other_user_id,
				message,
				created_at
			FROM messages
			WHERE sender_id = ? OR receiver_id = ?
			ORDER BY created_at DESC
		) m ON u.id = m.other_user_id
		WHERE u.id != ?
		GROUP BY u.id
		ORDER BY 
			CASE WHEN m.created_at IS NULL THEN 1 ELSE 0 END,
			m.created_at DESC,
			u.username ASC`,
		user_id, user_id, user_id, user_id)

	if err != nil {
		log.Printf("Error getting conversations: %v", err)
		return nil, err
	}
	defer rows.Close()

	var conversations []models.Conversation
	for rows.Next() {
		var conv models.Conversation
		var last_message, last_message_at sql.NullString
		if err := rows.Scan(&conv.User_id, &conv.Username, &conv.First_name, &conv.Last_name, &last_message, &last_message_at); err != nil {
			log.Printf("Error scanning conversation: %v", err)
			continue
		}

		if last_message.Valid {
			conv.Last_message = last_message.String
			conv.Last_message_at = last_message_at.String
		}

		conversations = append(conversations, conv)
	}

	return conversations, nil
}

func Get_username_by_id(db *sql.DB, user_id int) (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", user_id).Scan(&username)
	return username, err
}
