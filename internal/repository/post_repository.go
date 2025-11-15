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

func Insert_comment(db *sql.DB, content string, user_id, post_id int) error {
	_, err := db.Exec(`
		INSERT INTO comments (content, user_id, post_id) VALUES (?, ?, ?)`,
		content, user_id, post_id)
	return err
}

func Get_all_posts(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query(`SELECT p.id, p.title, p.content, p.category, u.username, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Username, &p.Created_at); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

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
			var c models.Comment
			if err := c_r.Scan(&c.Content, &c.Username, &c.Created_at); err != nil {
				return nil, err
			}
			comments = append(comments, c)
		}
		posts[i].Comments = comments
	}

	return posts, nil
}

func Get_post_by_id(db *sql.DB, post_id int) (*models.Post, error) {
	var p models.Post
	err := db.QueryRow(`
		SELECT p.id, p.title, p.content, p.category, u.username, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.id = ?`, post_id).Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Username, &p.Created_at)

	if err != nil {
		return nil, err
	}

	// Get comments for this post
	comments, err := Get_comments_by_post(db, post_id)
	if err != nil {
		return nil, err
	}
	p.Comments = comments

	return &p, nil
}

func Get_comments_by_post(db *sql.DB, post_id int) ([]models.Comment, error) {
	rows, err := db.Query(`
		SELECT c.content, u.username, c.created_at
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC`, post_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.Content, &c.Username, &c.Created_at); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
