package router

import (
	"database/sql"
	"net/http"
	"real_time_forum/internal/handler"
)

func New_Router(db *sql.DB) http.Handler {
	m := http.NewServeMux()

	auth_h := &handler.Auth_handler{DB: db}
	post_h := &handler.Post_handler{DB: db}
	chat_h := &handler.Chat_handler{DB: db}

	file_server := http.FileServer(http.Dir("./web/static"))
	m.Handle("/static/", http.StripPrefix("/static/", file_server))

	m.HandleFunc("/api/register", auth_h.Register)
	m.HandleFunc("/api/login", auth_h.Login)
	m.HandleFunc("/api/logout", auth_h.Logout)
	m.HandleFunc("/api/posts", post_h.Get_posts)

	m.HandleFunc("/api/posts/create", post_h.Auth_middleware(post_h.Create_post))
	m.HandleFunc("/api/comments/add", post_h.Auth_middleware(post_h.Create_comment))
	m.HandleFunc("/api/users", post_h.Auth_middleware(chat_h.Get_users))
	m.HandleFunc("/api/conversations", post_h.Auth_middleware(chat_h.Get_conversations))
	m.HandleFunc("/api/messages/", post_h.Auth_middleware(chat_h.Get_messages))
	m.HandleFunc("/api/posts/", post_h.Auth_middleware(func(w http.ResponseWriter, r *http.Request) {
		// Check if it's a request for comments
		if len(r.URL.Path) > len("/api/posts/") && r.URL.Path[len(r.URL.Path)-9:] == "/comments" {
			post_h.Get_post_comments(w, r)
		} else {
			post_h.Get_post_by_id(w, r)
		}
	}))

	m.HandleFunc("/ws", handler.Ws_handler(db))

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./web/templates/index.html")
	})

	return m
}
