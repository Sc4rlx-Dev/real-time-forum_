package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/repository"
	"strconv"
)

type Post_handler struct {
	DB *sql.DB
}

func (h *Post_handler) Auth_middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-User-ID", "1")
		next.ServeHTTP(w, r)
	}
}

func (h *Post_handler) Create_post(w http.ResponseWriter, r *http.Request) {
    var data map[string]string
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    user_id, _ := strconv.Atoi(r.Header.Get("X-User-ID"))
	err := repository.Insert_post(h.DB, data["title"], data["content"], data["category"], user_id)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
}

func (h *Post_handler) Create_comment(w http.ResponseWriter, r *http.Request) {
	content := "This is a great comment!"
	user_id := 1
	post_id := 1

	err := repository.Insert_comment(h.DB, content, user_id, post_id)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added successfully"})
}

func (h *Post_handler) Get_posts(w http.ResponseWriter, r *http.Request) {
	posts, err := repository.Get_all_posts(h.DB)
	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}