package models

type User_data struct {
    Username   string
    First_name string
    Last_name  string
    Age        int
    Email      string
    Password   string
    Gender     string
}

type Data struct {
    Username string 
    Password string
}

type Post struct {
    ID         int
    Title      string
    Content    string
    Category   string
    Username   string    
    Created_at string
    Comments   []Comment 
}

type Comment struct {
    Content    string
    Username   string 
    Created_at string
}
type Message struct {
    ID      string `json:"id"`
    Message string `json:"message"`
    From    string `json:"from"` 
    To      string `json:"to"`  
    Date    string `json:"date"` 
    Type    string `json:"type"` 
}