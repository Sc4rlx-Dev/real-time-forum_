package router

import (
    "database/sql"
    "net/http"
    "real_time_forum/internal/handler"
)

func NewRouter(db *sql.DB) http.Handler {
    m := http.NewServeMux()
    authHandler := &handler.AuthHandler{DB: db}
    m.HandleFunc("/register", authHandler.Register)
    m.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/posts/create", postHandler.authMiddleware(postHandler.CreatePost))
	mux.HandleFunc("/comments/add", postHandler.authMiddleware(postHandler.CreateComment))
	mux.HandleFunc("/posts", postHandler.GetPosts)

return m
}