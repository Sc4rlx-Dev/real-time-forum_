package router

import (
    "database/sql"
    "net/http"
    "real_time_forum/internal/handler"
)

func New_router(db *sql.DB) http.Handler {
    m := http.NewServeMux()

    auth_h := &handler.Auth_handler{DB: db}
    post_h := &handler.Post_handler{DB: db}

    m.HandleFunc("/register", auth_h.Register)
    m.HandleFunc("/login", auth_h.Login)
    m.HandleFunc("/posts/create", post_h.Auth_middleware(post_h.Create_post))
    m.HandleFunc("/comments/add", post_h.Auth_middleware(post_h.Create_comment))
    m.HandleFunc("/posts", post_h.Get_posts)

    return m
}