# Real-Time Forum

A modern forum application with real-time chat capabilities built with Go and vanilla JavaScript.

## Features

- **User Authentication**: Registration and login with secure session management
- **Forum Posts**: Create, view posts with categories
- **Comments**: Add comments to posts
- **Real-Time Chat**: WebSocket-based instant messaging between users
- **Online Status**: See who's currently online
- **Responsive UI**: Clean, modern interface

## Technology Stack

### Backend
- **Go 1.24+**: Modern, efficient backend
- **SQLite3**: Lightweight database
- **Gorilla WebSocket**: Real-time communication
- **bcrypt**: Secure password hashing
- **UUID**: Session token generation

### Frontend
- **Vanilla JavaScript (ES6 Modules)**: No framework dependencies
- **WebSocket API**: Real-time updates
- **Modern CSS**: Clean styling

## Project Structure

```
real-time-forum/
├── cmd/
│   └── web/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/               # Configuration (future)
│   ├── handler/              # HTTP and WebSocket handlers
│   │   ├── auth_handler.go   # Login/Register endpoints
│   │   ├── post_handler.go   # Post/Comment endpoints
│   │   ├── chat_handler.go   # Chat API endpoints
│   │   └── websocket_handler.go # WebSocket connection
│   ├── models/               # Data structures
│   │   └── models.go
│   ├── repository/           # Database layer
│   │   ├── database.go       # DB initialization
│   │   ├── user_repository.go
│   │   ├── post_repository.go
│   │   └── chat_repository.go
│   ├── router/               # HTTP routing
│   │   └── router.go
│   └── service/              # Business logic (future)
├── migrations/
│   └── test.sql              # Database schema
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css     # Application styles
│   │   └── js/
│   │       ├── app.js        # Main application logic
│   │       ├── api.js        # API calls
│   │       ├── ui.js         # UI rendering
│   │       └── websocket.js  # WebSocket client
│   └── templates/
│       └── index.html        # Single-page application
└── storage/                  # SQLite database location
```

## Getting Started

### Prerequisites

- Go 1.24 or higher
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Sc4rlx-Dev/real-time-forum.git
cd real-time-forum
```

2. Install Go dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o forum ./cmd/web/main.go
```

### Running the Application

1. Start the server:
```bash
./forum
```
Or use `go run`:
```bash
go run ./cmd/web/main.go
```

2. Open your browser and navigate to:
```
http://localhost:8081
```

### Database

The application automatically:
- Creates the `storage` directory
- Initializes the SQLite database
- Sets up all required tables from `migrations/test.sql`

## API Endpoints

### Public Endpoints
- `GET /` - Serve the main application
- `GET /static/*` - Serve static assets
- `POST /api/register` - Register a new user
- `POST /api/login` - Login user

### Protected Endpoints (require authentication)
- `POST /api/logout` - Logout and clear session
- `GET /api/posts` - Get all posts with comments
- `GET /api/posts/{id}` - Get a single post by ID
- `GET /api/posts/{id}/comments` - Get comments for a specific post
- `POST /api/posts/create` - Create a new post
- `POST /api/comments/add` - Add a comment to a post
- `GET /api/users` - Get all users (for chat)
- `GET /api/conversations` - Get conversation list with last messages
- `GET /api/messages/{username}?before=&limit=` - Get paginated chat history
- `GET /ws` - WebSocket connection for real-time chat

## Database Schema

### Users Table
- Stores user information with hashed passwords
- Fields: id, username, email, password, age, firstname, lastname, gender, created_at

### Sessions Table
- Manages user sessions with expiry
- Fields: id, user_id, username, session_id, expiry_date
- Sessions expire after 24 hours

### Categories Table
- Predefined post categories
- Fields: id, name
- Default categories: General, Technology, Sports, Entertainment, News, Gaming, Music, Other

### Posts Table
- Forum posts with category support
- Fields: id, title, content, category, user_id, created_at
- Foreign keys: user_id references users(id) with CASCADE delete

### Comments Table
- Comments on posts
- Fields: id, content, user_id, post_id, created_at
- Foreign keys: user_id and post_id with CASCADE delete

### Conversations & Messages Tables
- Real-time chat support
- Conversations: id, sender_id, receiver_id, created_at
- Messages: id, message, conversation_id, sender_id, receiver_id, created_at
- Maintains full message history between users

## Development

### Code Formatting
```bash
go fmt ./...
```

### Code Vetting
```bash
go vet ./...
```

### Tidy Dependencies
```bash
make tidy
# or
go mod tidy
```

## Features in Detail

### Authentication
- Secure password hashing with bcrypt (default cost)
- Session-based authentication with UUID v4 tokens
- 24-hour automatic session expiry
- HttpOnly cookies for session token (XSS protection)
- Comprehensive input validation:
  - Username: 3-20 characters, alphanumeric + underscore
  - Email: Valid email format with regex validation
  - Password: Minimum 6 characters
  - Age: Range 13-120 years
  - Names: 2-50 characters
- Field-level error messages
- Unique constraint enforcement (username, email)
- Login accepts username OR email
- Proper logout with session deletion

### Real-Time Chat
- WebSocket-based instant messaging
- Online/offline user status tracking
- Message persistence in database
- Automatic reconnection handling
- Conversation list with last message preview
- Messages sorted by time (newest first)
- Pagination support (10 messages per page)
- Offline message queuing
- Real-time user list updates

### Forum
- Create posts with categories
- View all posts with comments
- Get single post by ID
- Get comments for specific post
- Real-time updates when logged in
- User attribution for all content
- Chronological comment ordering

## Security Considerations

- Passwords are hashed using bcrypt
- Session tokens are UUID v4 for uniqueness
- Session expiry enforced at database level
- WebSocket connections require valid session
- SQL injection prevention via prepared statements

## Contributing

Contributions are welcome! Please follow these guidelines:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Authors

- Sc4rlx-Dev

## Acknowledgments

- Built with Go and vanilla JavaScript
- Uses Gorilla WebSocket for real-time features
- SQLite for simple, embedded database
