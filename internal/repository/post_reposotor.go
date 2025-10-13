package repository

import (
	"database/sql"
	"real_time_forum/internal/models"
)

func InsertPost(db *sql.DB, title, content, category string, user_id int) error {
	_, err := db.Exec(
		`INSERT INTO posts (title, content, category, user_id) VALUES (?, ?, ?, ?)`,
		title, content, category, user_id)
	return err
}

func InsertComment(db *sql.DB, content string, user_id, post_id int) error {
	_, err := db.Exec(`
		INSERT INTO comments (content, user_id, post_id) VALUES (?, ?, ?)`,
		content, user_id, post_id)
	return err
}

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
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
	//for each post -> fetch its commants
	for i := range posts {
		c_r, err := db.Query(`
			SELECT c.content, u.username, c.created_at
			FROM comments c
			JOIN users u ON c.user_id = u.id
			WHERE c.post_id = ?
			ORDER BY c.created_at ASC`, posts[i].ID)
		if err != nil {
			return nil, err
		}
		defer c_r.Close()

		var comments []models.Comment
		for c_r.Next() {
			var comment models.Comment
			if err := c_r.Scan(&comment.Content, &comment.Username, &comment.CreatedAt); err != nil {
				return nil, err
			}
			comments = append(comments, comment)
		}
		posts[i].Comments = comments
	}

	return posts, nil	
}