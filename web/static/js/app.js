import { 
    render_login_form, 
    render_home_page, 
    render_posts,
    render_user_list,
    render_chat_window,
    append_message,
    render_register_form,
    render_create_post_modal,
    show_notification,
    show_loading,
    hide_loading
} from './ui.js';
import { 
    check_auth_status, 
    login_user, 
    get_posts,
    get_users,
    get_chat_history,
    get_cookie,
    register_user,
    create_post,
    add_comment,
    logout_user
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
    render_home_page(current_username, handle_logout, handle_create_post);
    
    show_loading();
    const posts = await get_posts();
    hide_loading();
    render_posts(posts, handle_add_comment);

    all_users = await get_users();
    render_user_list(all_users, online_users, on_user_click);

    connect_websocket(current_username);

    set_on_message(handle_incoming_message);
    set_on_user_list(handle_user_list_update);
}

function show_login_page() {
    render_login_form(async (form_data) => {
        const result = await login_user(form_data);
        if (result.success) {
            current_username = result.username;
            show_notification('Login successful!', 'success');
            show_home_page();
        } else {
            show_notification('Login failed! Please check your credentials.', 'error');
        }
    }, () => {
        show_register_page();
    });
}

function show_register_page() {
    render_register_form(async (form_data) => {
        const success = await register_user(form_data);
        if (success) {
            show_notification('Registration successful! Please log in.', 'success');
            show_login_page();
        } else {
            show_notification('Registration failed. Username or email may already exist.', 'error');
        }
    }, () => {
        show_login_page();
    });
}

async function handle_logout() {
    await logout_user();
    show_notification('Logged out successfully', 'info');
    window.location.reload();
}

function handle_create_post() {
    render_create_post_modal(async (post_data) => {
        show_loading();
        const success = await create_post(post_data);
        hide_loading();
        
        if (success) {
            show_notification('Post created successfully!', 'success');
            const posts = await get_posts();
            render_posts(posts, handle_add_comment);
        } else {
            show_notification('Failed to create post. Please try again.', 'error');
        }
    }, () => {});
}

async function handle_add_comment(post_id, content, post_element) {
    const success = await add_comment(post_id, content);
    
    if (success) {
        show_notification('Comment added successfully!', 'success');
        const posts = await get_posts();
        render_posts(posts, handle_add_comment);
    } else {
        show_notification('Failed to add comment. Please try again.', 'error');
    }
}

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
        show_notification(`New message from ${msg.From_username}`, 'info');
    }
}

function handle_user_list_update(new_online_users) {
    online_users = new_online_users;
    render_user_list(all_users, online_users, on_user_click);
}

main();