package router

import (
    "database/sql"
    "net/http"
    "real_time_forum/internal/handler"
)

func NewRouter(db *sql.DB) http.Handler {
    m := http.NewServeMux()
    authHandler := &handler.AuthHandler{DB: db}
    postHandler := &handler.PostHandler{DB: db}
    m.HandleFunc("/register", authHandler.Register)
    m.HandleFunc("/login", auth.Login)
    m.HandleFunc("/posts/create", postHandler.authMiddleware(postHandler.CreatePost))
    m.HandleFunc("/comments/add", postHandler.authMiddleware(postHandler.CreateComment))
    m.HandleFunc("/posts", postHandler.GetPosts)

return m
}