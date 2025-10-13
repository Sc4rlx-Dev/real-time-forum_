package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/repository"
)

type PostHandler struct {
	DB *sql.DB
}

func (h *PostHandler) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-User-ID", "1")
		next.ServeHTTP(w, r)
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	title := "test post"
	content := "hhhhhhhhhhhhhhhhhhhhhhh"
	category := "test category"
	userID := 1

	err := repository.InsertPost(h.DB, title, content, category, userID)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
}

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	content := "This is a great comment!"
	userID := 1
	postID := 1

	err := repository.InsertComment(h.DB, content, userID, postID)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := repository.GetAllPosts(h.DB)
	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
