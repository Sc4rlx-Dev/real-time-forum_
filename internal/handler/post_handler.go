package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"real_time_forum/internal/repository"
	"strconv"
	"strings"
)

type Post_handler struct {
	DB *sql.DB
}

// UPDATED: Real Authentication Middleware
func (h *Post_handler) Auth_middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized: No session cookie", http.StatusUnauthorized)
			return
		}

		user_id, _, err := repository.Get_user_from_session(h.DB, cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
			return
		}

		// Add the user ID to the request header for other handlers to use
		r.Header.Set("X-User-ID", strconv.Itoa(user_id))
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
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user_id, _ := strconv.Atoi(r.Header.Get("X-User-ID"))
	post_id, _ := strconv.Atoi(data["post_id"])

	err := repository.Insert_comment(h.DB, data["content"], user_id, post_id)
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

func (h *Post_handler) Get_post_by_id(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL path
	post_id_str := r.URL.Path[len("/api/posts/"):]
	post_id, err := strconv.Atoi(post_id_str)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := repository.Get_post_by_id(h.DB, post_id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *Post_handler) Get_post_comments(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	post_id, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	comments, err := repository.Get_comments_by_post(h.DB, post_id)
	if err != nil {
		http.Error(w, "Failed to retrieve comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
