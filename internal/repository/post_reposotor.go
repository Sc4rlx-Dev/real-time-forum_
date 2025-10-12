package repository

import (
	"database/sql"
	"real_time_forum/internal/models"
)

func Insert_post(db *sql.DB, title, content, category string, user_id int) error {
	_, err := db.Exec(
		`INSERT INTO posts (title, content, category, user_id) VALUES (?, ?, ?, ?)`,
		title, content, category, user_id)
	return err
}

func Insert_commant(db *sql.DB, content string, user_id, post_id int) error {
	_, err := db.Exec(`
		INSERT INTO comments (content, user_id, post_id) VALUES (?, ?, ?)`,
		content, user_id, post_id)
	return err
}

func get_all_posts(db *sql.DB) ([]models.Post, error) {
	rows,err:= db.Query(`SELECT p.id, p.title, p.content, p.category, u.username, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC`)
	if err != nil { return nil, err }
	defer rows.Close()

	var posts []models
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Username, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
}