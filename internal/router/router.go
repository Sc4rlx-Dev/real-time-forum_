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

    file_server := http.FileServer(http.Dir("./web/static"))
    m.Handle("/static/", http.StripPrefix("/static/", file_server))

    m.HandleFunc("/api/register", auth_h.Register)
    m.HandleFunc("/api/login", auth_h.Login)
    m.HandleFunc("/api/posts/create", post_h.Auth_middleware(post_h.Create_post))
    m.HandleFunc("/api/comments/add", post_h.Auth_middleware(post_h.Create_comment))
    m.HandleFunc("/api/posts", post_h.Get_posts)

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
