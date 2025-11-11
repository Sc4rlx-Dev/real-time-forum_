// --- HELPER FUNCTION ---
function create_element(tag, className, text) {
    const el = document.createElement(tag);
    if (className) el.className = className;
    if (text) el.textContent = text;
    return el;
}

// --- AUTH FORMS ---
export function render_login_form(on_login, on_switch_to_register) {
    document.body.innerHTML = "";
    const c = create_element('div', 'form-container');
    const title = create_element('h1', '', 'Login');
    const form = create_element('form');
    
    const username_inpt = create_element('input');
    username_inpt.placeholder = 'Username or Email';
    username_inpt.name = 'username';

    const passwd_inpt = create_element('input');
    passwd_inpt.type = 'password';
    passwd_inpt.placeholder = 'Password';
    passwd_inpt.name = 'password';

    const sub_button = create_element('button', '', 'Login');
    sub_button.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const form_data = new FormData(form);
        on_login(form_data);
    });

    const togl_txt = create_element('p', 'toggle-form', "Don't have an account? Register");
    togl_txt.onclick = on_switch_to_register;

    form.append(username_inpt, passwd_inpt, sub_button);
    c.append(title, form, togl_txt);
    document.body.append(c);
}

// --- NEW FUNCTION ---
// Renders the registration form
export function render_register_form(on_register, on_switch_to_login) {
    document.body.innerHTML = "";
    const c = create_element('div', 'form-container');
    const title = create_element('h1', '', 'Register');
    const form = create_element('form');

    // Create inputs for all fields required by the backend
    const username_inpt = create_element('input');
    username_inpt.placeholder = 'Username';
    username_inpt.name = 'Username';

    const email_inpt = create_element('input');
    email_inpt.type = 'email';
    email_inpt.placeholder = 'Email';
    email_inpt.name = 'Email';

    const passwd_inpt = create_element('input');
    passwd_inpt.type = 'password';
    passwd_inpt.placeholder = 'Password';
    passwd_inpt.name = 'Password';

    const first_name_inpt = create_element('input');
    first_name_inpt.placeholder = 'First Name';
    first_name_inpt.name = 'First_name';

    const last_name_inpt = create_element('input');
    last_name_inpt.placeholder = 'Last Name';
    last_name_inpt.name = 'Last_name';

    const age_inpt = create_element('input');
    age_inpt.type = 'number';
    age_inpt.placeholder = 'Age';
    age_inpt.name = 'Age';

    // Create a select dropdown for Gender
    const gender_select = create_element('select');
    gender_select.name = 'Gender';
    const option_male = create_element('option', '', 'Male');
    option_male.value = 'Male';
    const option_female = create_element('option', '', 'Female');
    option_female.value = 'Female';
    gender_select.append(option_male, option_female);

    const sub_button = create_element('button', '', 'Register');
    sub_button.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const form_data = new FormData(form);
        on_register(form_data);
    });

    const togl_txt = create_element('p', 'toggle-form', "Already have an account? Login");
    togl_txt.onclick = on_switch_to_login;

    form.append(
        username_inpt, 
        email_inpt, 
        passwd_inpt, 
        first_name_inpt, 
        last_name_inpt, 
        age_inpt, 
        gender_select, 
        sub_button
    );
    c.append(title, form, togl_txt);
    document.body.append(c);
}
// --------------------


// --- HOME PAGE (Updated) ---
export function render_home_page() {
    document.body.innerHTML = '';
    document.body.className = 'home-layout';

    const home_container = create_element('div', 'home-container');
    const posts_column = create_element('div', 'posts-column');
    const chat_column = create_element('div', 'chat-column');

    const page_title = create_element('h1', 'page-title', 'Forum Posts');
    const posts_container = create_element('div', 'posts-list');
    posts_container.id = 'posts-container';
    
    const chat_title = create_element('h1', 'page-title', 'Chat');
    const user_list = create_element('div', 'user-list');
    user_list.id = 'user-list';
    
    const chat_window = create_element('div', 'chat-window');
    chat_window.id = 'chat-window';
    chat_window.style.display = 'none'; // Hide until a user is clicked

    posts_column.append(page_title, posts_container);
    chat_column.append(chat_title, user_list, chat_window);
    home_container.append(posts_column, chat_column);
    document.body.append(home_container);
}

// --- POSTS (Unchanged) ---
export function render_posts(posts_data) {
    const posts_container = document.getElementById('posts-container');
    if (!posts_container) return;
    posts_container.innerHTML = '';
    if (!posts_data || posts_data.length === 0) {
        posts_container.append(create_element('p', 'no-posts', 'No posts yet!'));
        return;
    }
    posts_data.forEach(post => {
        const post_element = create_element('article', 'post');
        const title = create_element('h2', 'post-title', post.Title);
        const meta = create_element('p', 'post-meta', `Posted by ${post.Username} in ${post.Category} on ${post.Created_at}`);
        const content = create_element('p', 'post-content', post.Content);
        const comments_section = create_element('div', 'comments-section');
        const comments_title = create_element('h3', 'comments-title', 'Comments');
        comments_section.append(comments_title);
        if (post.Comments && post.Comments.length > 0) {
            post.Comments.forEach(comment => {
                const comment_element = create_element('div', 'comment');
                const comment_meta = create_element('p', 'comment-meta', `Comment by ${comment.Username} on ${comment.Created_at}`);
                const comment_content = create_element('p', 'comment-content', comment.Content);
                comment_element.append(comment_meta, comment_content);
                comments_section.append(comment_element);
            });
        } else {
            comments_section.append(create_element('p', 'no-comments', 'No comments yet.'));
        }
        post_element.append(title, meta, content, comments_section);
        posts_container.append(post_element);
    });
}

// --- NEW CHAT FUNCTIONS ---
export function render_user_list(users, online_users, on_user_click) {
    const user_list_el = document.getElementById('user-list');
    if (!user_list_el) return;
    user_list_el.innerHTML = '';

    users.forEach(user => {
        const is_online = online_users.includes(user.Username); // Corrected: user.Username
        const user_item = create_element('div', 'user-item');
        user_item.textContent = user.Username; // Corrected: user.Username
        
        const status_indicator = create_element('span', is_online ? 'online' : 'offline');
        user_item.append(status_indicator);
        
        user_item.onclick = () => on_user_click(user.Username); // Corrected: user.Username
        user_list_el.append(user_item);
    });
}

export function render_chat_window(username, messages, on_send_message) {
    const chat_window = document.getElementById('chat-window');
    chat_window.innerHTML = '';
    chat_window.style.display = 'flex';

    const chat_header = create_element('div', 'chat-header', `Chat with ${username}`);
    const message_list = create_element('div', 'message-list');
    message_list.id = `messages-${username}`;

    messages.forEach(msg => {
        append_message(msg);
    });

    const send_form = create_element('form', 'send-form');
    const message_input = create_element('input');
    message_input.placeholder = 'Type a message...';
    const send_button = create_element('button', '', 'Send');
    send_button.type = 'submit';

    send_form.addEventListener('submit', (e) => {
        e.preventDefault();
        const message_text = message_input.value;
        if (message_text.trim()) {
            on_send_message(message_text);
            message_input.value = '';
        }
    });

    send_form.append(message_input, send_button);
    chat_window.append(chat_header, message_list, send_form);
}

export function append_message(msg, current_username) {
    const other_user = msg.From_username === current_username ? msg.To_username : msg.From_username;
    const message_list = document.getElementById(`messages-${other_user}`);
    
    if (message_list) {
        const msg_item = create_element('div', 'message-item');
        
        if (msg.From_username === current_username) {
            msg_item.classList.add('sent');
        } else {
            msg_item.classList.add('received');
        }
        
        msg_item.textContent = msg.Message; // Corrected: msg.Message
        message_list.append(msg_item);
        message_list.scrollTop = message_list.scrollHeight;
    }
}