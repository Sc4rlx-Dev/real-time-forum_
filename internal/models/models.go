package models

type User_data struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        int    `json:"-"` // Don't send age to frontend
	Email      string `json:"-"` // Don't send email
	Password   string `json:"-"` // Never send password
	Gender     string `json:"-"`
}

type Data struct {
	Username string
	Password string
}

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   string    `json:"category"`
	Username   string    `json:"username"`
	Created_at string    `json:"created_at"`
	Comments   []Comment `json:"comments"`
}

type Comment struct {
	Content    string `json:"content"`
	Username   string `json:"username"`
	Created_at string `json:"created_at"`
}

type Message struct {
	ID            string      `json:"id"`
	Message       string      `json:"message"`
	From_username string      `json:"from_username"`
	To_username   string      `json:"to_username"`
	Date          string      `json:"date"`
	Type          string      `json:"type"`
	Data          interface{} `json:"data,omitempty"` // For user lists, etc.
}

type Conversation struct {
	User_id         int    `json:"user_id"`
	Username        string `json:"username"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Last_message    string `json:"last_message"`
	Last_message_at string `json:"last_message_at"`
}
