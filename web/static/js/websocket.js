let socket = null;
let current_username = null;

// Callbacks for the main app to use
let on_message_received = null;
let on_user_list_updated = null;

export function connect_websocket(username) {
    if (socket) {
        return; // Already connected
    }
    
    current_username = username;
    
    // Connect using the session cookie for auth
    socket = new WebSocket(`ws://${window.location.host}/ws`);

    socket.onopen = () => {
        console.log("WebSocket connected.");
    };

    socket.onmessage = (event) => {
        const msg = JSON.parse(event.data);

        if (msg.type === 'user_list') {
            if (on_user_list_updated) {
                on_user_list_updated(msg.data);
            }
        } else {
            if (on_message_received) {
                on_message_received(msg);
            }
        }
    };

    socket.onclose = () => {
        console.log("WebSocket disconnected.");
        socket = null;
        // Optionally try to reconnect
    };

    socket.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

export function send_chat_message(to_username, message_text) {
    if (!socket || socket.readyState !== WebSocket.OPEN) {
        console.error("WebSocket is not connected.");
        return;
    }

    const message = {
        id: "", // Will be set by backend
        message: message_text,
        from_username: current_username,
        to_username: to_username,
        date: new Date().toISOString(),
        type: "chat_message",
    };

    socket.send(JSON.stringify(message));
}

// Setters for callbacks
export function set_on_message(callback) {
    on_message_received = callback;
}
export function set_on_user_list(callback) {
    on_user_list_updated = callback;
}
