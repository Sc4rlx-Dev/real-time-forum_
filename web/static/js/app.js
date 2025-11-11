import { 
    render_login_form, 
    render_home_page, 
    render_posts,
    render_user_list,
    render_chat_window,
    append_message,
    render_register_form // --- ADDED
} from './ui.js';
import { 
    check_auth_status, 
    login_user, 
    get_posts,
    get_users,
    get_chat_history,
    get_cookie,
    register_user // --- ADDED
} from './api.js';
import {
    connect_websocket,
    send_chat_message,
    set_on_message,
    set_on_user_list
} from './websocket.js';

let all_users = [];
let online_users = [];
let current_username = null;
let current_chat_with = null;

async function main() {
    const is_logged_in = await check_auth_status();

    if (is_logged_in) {
        current_username = get_cookie("username"); 
        if (current_username) {
            show_home_page()
        } else {
            show_login_page()
        }
    } else {
        show_login_page()
    }
}

async function show_home_page() {
    render_home_page();
    
    const posts = await get_posts();
    render_posts(posts);

    all_users = await get_users();
    render_user_list(all_users, online_users, on_user_click);

    connect_websocket(current_username);

    set_on_message(handle_incoming_message);
    set_on_user_list(handle_user_list_update);
}

function show_login_page() {
    render_login_form(async (form_data) => {
        const success = await login_user(form_data);
        if (success) {
            // --- THIS IS THE FIX ---
            // Read the "username" cookie, which is visible to JavaScript.
            current_username = get_cookie("session_token");
            // ---------------------
            console.log("currrrrr",current_username , typeof current_username)
            if (current_username) {
                show_home_page();
            } else {
                // This alert should no longer appear
                alert('Login successful, but cookie was not set.');
            }
        } else {
            alert('Login failed! Please try again.');
        }
    }, () => {
        // Switch to the register page instead of showing an alert
        show_register_page();
        // ---------------
    });
}

// --- NEW FUNCTION ---
// Renders the register page and handles form submission
function show_register_page() {
    render_register_form(async (form_data) => {
        // This is the on_register callback
        const success = await register_user(form_data);
        if (success) {
            alert('Registration successful! Please log in.');
            show_login_page(); // Go back to login page
        } else {
            alert('Registration failed. Please try again.');
        }
    }, () => {
        // This is the on_switch_to_login callback
        show_login_page(); // Go back to login page
    });
}
// --------------------


// --- CHAT LOGIC ---

async function on_user_click(username) {
    current_chat_with = username;
    const history = await get_chat_history(username);
    render_chat_window(username, history, (message_text) => {
        send_chat_message(current_chat_with, message_text);
        
        append_message({
            From_username: current_username,
            To_username: current_chat_with,
            Message: message_text
        }, current_username);
    });
}

function handle_incoming_message(msg) {
    if (msg.From_username === current_chat_with) {
        append_message(msg, current_username);
    } else {
        console.log(`Received message from ${msg.From_username}, but not in active chat.`);
    }
}

function handle_user_list_update(new_online_users) {
    online_users = new_online_users;
    render_user_list(all_users, online_users, on_user_click);
}

// Start the application.
main();