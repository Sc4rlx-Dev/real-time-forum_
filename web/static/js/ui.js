function create_element(tag, className, text) {
    const el = document.createElement(tag);
    if (className) el.className = className;
    if (text) el.textContent = text;
    return el;
}

export function render_login_form(on_login, on_switch_to_register) {
    document.body.innerHTML = "";
    const c = create_element('div', 'form-container');
    const title = create_element('h1', '', 'Login');
    const form = create_element('form');
    
    const username_inpt = create_element('input');
    username_inpt.placeholder = 'Username or Email';
    username_inpt.name = 'username';
    username_inpt.required = true;
    username_inpt.minLength = 3;

    const passwd_inpt = create_element('input');
    passwd_inpt.type = 'password';
    passwd_inpt.placeholder = 'Password';
    passwd_inpt.name = 'password';
    passwd_inpt.required = true;
    passwd_inpt.minLength = 6;

    const sub_button = create_element('button', '', 'Login');
    sub_button.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        if (!form.checkValidity()) {
            return;
        }
        const form_data = new FormData(form);
        sub_button.disabled = true;
        sub_button.textContent = 'Logging in...';
        on_login(form_data).finally(() => {
            sub_button.disabled = false;
            sub_button.textContent = 'Login';
        });
    });

    const togl_txt = create_element('p', 'toggle-form', "Don't have an account? Register");
    togl_txt.onclick = on_switch_to_register;

    form.append(username_inpt, passwd_inpt, sub_button);
    c.append(title, form, togl_txt);
    document.body.append(c);
}

export function render_register_form(on_register, on_switch_to_login) {
    document.body.innerHTML = "";
    const c = create_element('div', 'form-container');
    const title = create_element('h1', '', 'Register');
    const form = create_element('form');

    const username_inpt = create_element('input');
    username_inpt.placeholder = 'Username';
    username_inpt.name = 'Username';
    username_inpt.required = true;
    username_inpt.minLength = 3;
    username_inpt.maxLength = 20;
    username_inpt.pattern = '[a-zA-Z0-9_]+';
    username_inpt.title = 'Username must be 3-20 characters and contain only letters, numbers, and underscores';

    const email_inpt = create_element('input');
    email_inpt.type = 'email';
    email_inpt.placeholder = 'Email';
    email_inpt.name = 'Email';
    email_inpt.required = true;

    const passwd_inpt = create_element('input');
    passwd_inpt.type = 'password';
    passwd_inpt.placeholder = 'Password';
    passwd_inpt.name = 'Password';
    passwd_inpt.required = true;
    passwd_inpt.minLength = 6;

    const first_name_inpt = create_element('input');
    first_name_inpt.placeholder = 'First Name';
    first_name_inpt.name = 'First_name';
    first_name_inpt.required = true;
    first_name_inpt.minLength = 2;

    const last_name_inpt = create_element('input');
    last_name_inpt.placeholder = 'Last Name';
    last_name_inpt.name = 'Last_name';
    last_name_inpt.required = true;
    last_name_inpt.minLength = 2;

    const age_inpt = create_element('input');
    age_inpt.type = 'number';
    age_inpt.placeholder = 'Age';
    age_inpt.name = 'Age';
    age_inpt.required = true;
    age_inpt.min = 13;
    age_inpt.max = 120;

    const gender_select = create_element('select');
    gender_select.name = 'Gender';
    gender_select.required = true;
    const option_placeholder = create_element('option', '', 'Select Gender');
    option_placeholder.value = '';
    option_placeholder.disabled = true;
    option_placeholder.selected = true;
    const option_male = create_element('option', '', 'Male');
    option_male.value = 'Male';
    const option_female = create_element('option', '', 'Female');
    option_female.value = 'Female';
    gender_select.append(option_placeholder, option_male, option_female);

    const sub_button = create_element('button', '', 'Register');
    sub_button.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        if (!form.checkValidity()) {
            return;
        }
        const form_data = new FormData(form);
        sub_button.disabled = true;
        sub_button.textContent = 'Registering...';
        on_register(form_data).finally(() => {
            sub_button.disabled = false;
            sub_button.textContent = 'Register';
        });
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

export function render_home_page(username, on_logout, on_create_post) {
    document.body.innerHTML = '';
    document.body.className = 'home-layout';

    const header = create_element('header', 'header');
    const header_content = create_element('div', 'header-content');
    const logo = create_element('h1', 'logo', 'Real-Time Forum');
    const user_info = create_element('div', 'user-info');
    const welcome_text = create_element('span', 'welcome-text', `Welcome, ${username}`);
    const logout_btn = create_element('button', 'logout-btn', 'Logout');
    logout_btn.onclick = on_logout;
    
    user_info.append(welcome_text, logout_btn);
    header_content.append(logo, user_info);
    header.append(header_content);

    const home_container = create_element('div', 'home-container');
    const posts_column = create_element('div', 'posts-column');
    const chat_column = create_element('div', 'chat-column');

    const posts_header = create_element('div', 'posts-header');
    const page_title = create_element('h1', 'page-title', 'Forum Posts');
    const create_post_btn = create_element('button', 'create-post-btn', '+ Create Post');
    create_post_btn.onclick = on_create_post;
    posts_header.append(page_title, create_post_btn);

    const posts_container = create_element('div', 'posts-list');
    posts_container.id = 'posts-container';
    
    const chat_title = create_element('h1', 'page-title', 'Chat');
    const user_list = create_element('div', 'user-list');
    user_list.id = 'user-list';
    
    const chat_window = create_element('div', 'chat-window');
    chat_window.id = 'chat-window';
    chat_window.style.display = 'none';

    posts_column.append(posts_header, posts_container);
    chat_column.append(chat_title, user_list, chat_window);
    home_container.append(posts_column, chat_column);
    document.body.append(header, home_container);
}

export function render_posts(posts_data, on_add_comment) {
    const posts_container = document.getElementById('posts-container');
    if (!posts_container) return;
    posts_container.innerHTML = '';
    if (!posts_data || posts_data.length === 0) {
        posts_container.append(create_element('p', 'no-posts', 'No posts yet! Create the first one.'));
        return;
    }
    posts_data.forEach(post => {
        const post_element = create_element('article', 'post');
        const title = create_element('h2', 'post-title', post.title || post.Title);
        const meta = create_element('p', 'post-meta', `Posted by ${post.username || post.Username} in ${post.category || post.Category} on ${new Date(post.created_at || post.Created_at).toLocaleString()}`);
        const content = create_element('p', 'post-content', post.content || post.Content);
        const comments_section = create_element('div', 'comments-section');
        const comments_title = create_element('h3', 'comments-title', 'Comments');
        comments_section.append(comments_title);
        
        const comments = post.comments || post.Comments || [];
        if (comments.length > 0) {
            comments.forEach(comment => {
                const comment_element = create_element('div', 'comment');
                const comment_meta = create_element('p', 'comment-meta', `Comment by ${comment.username || comment.Username} on ${new Date(comment.created_at || comment.Created_at).toLocaleString()}`);
                const comment_content = create_element('p', 'comment-content', comment.content || comment.Content);
                comment_element.append(comment_meta, comment_content);
                comments_section.append(comment_element);
            });
        } else {
            comments_section.append(create_element('p', 'no-comments', 'No comments yet.'));
        }

        const comment_form = create_element('form', 'comment-form');
        const comment_input = create_element('textarea', 'comment-input');
        comment_input.placeholder = 'Add a comment...';
        comment_input.rows = 2;
        const comment_btn = create_element('button', 'comment-btn', 'Add Comment');
        comment_btn.type = 'submit';
        
        comment_form.addEventListener('submit', (e) => {
            e.preventDefault();
            const comment_text = comment_input.value.trim();
            if (comment_text) {
                on_add_comment(post.id || post.ID, comment_text, post_element);
                comment_input.value = '';
            }
        });
        
        comment_form.append(comment_input, comment_btn);
        comments_section.append(comment_form);
        post_element.append(title, meta, content, comments_section);
        posts_container.append(post_element);
    });
}

export function render_user_list(users, online_users, on_user_click) {
    const user_list_el = document.getElementById('user-list');
    if (!user_list_el) return;
    user_list_el.innerHTML = '';

    users.forEach(user => {
        const is_online = online_users.includes(user.Username);
        const user_item = create_element('div', 'user-item');
        user_item.textContent = user.Username;
        
        const status_indicator = create_element('span', is_online ? 'online' : 'offline');
        user_item.append(status_indicator);
        
        user_item.onclick = () => on_user_click(user.Username);
        user_list_el.append(user_item);
    });
}

export function render_conversation_list(conversations, online_users, on_user_click) {
    const user_list_el = document.getElementById('user-list');
    if (!user_list_el) return;
    user_list_el.innerHTML = '';

    conversations.forEach(conv => {
        const is_online = online_users.includes(conv.username);
        const user_item = create_element('div', 'user-item');
        
        const user_info = create_element('div', 'user-info');
        const username_el = create_element('div', 'username', conv.username);
        user_info.appendChild(username_el);
        
        if (conv.last_message) {
            const last_msg = create_element('div', 'last-message', 
                conv.last_message.length > 30 ? conv.last_message.substring(0, 30) + '...' : conv.last_message);
            user_info.appendChild(last_msg);
        }
        
        const status_indicator = create_element('span', is_online ? 'online' : 'offline');
        user_item.appendChild(user_info);
        user_item.appendChild(status_indicator);
        
        user_item.onclick = () => on_user_click(conv.username);
        user_list_el.appendChild(user_item);
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
        
        msg_item.textContent = msg.Message;
        message_list.append(msg_item);
        message_list.scrollTop = message_list.scrollHeight;
    }
}

export function render_create_post_modal(on_submit, on_cancel) {
    const modal = create_element('div', 'modal');
    const modal_content = create_element('div', 'modal-content');
    const modal_header = create_element('h2', 'modal-header', 'Create New Post');
    const form = create_element('form', 'create-post-form');

    const title_input = create_element('input', 'post-title-input');
    title_input.placeholder = 'Post Title';
    title_input.name = 'title';
    title_input.required = true;

    const category_input = create_element('input', 'post-category-input');
    category_input.placeholder = 'Category (e.g., General, Tech, Sports)';
    category_input.name = 'category';
    category_input.required = true;

    const content_textarea = create_element('textarea', 'post-content-input');
    content_textarea.placeholder = 'What do you want to share?';
    content_textarea.name = 'content';
    content_textarea.rows = 6;
    content_textarea.required = true;

    const button_group = create_element('div', 'button-group');
    const cancel_btn = create_element('button', 'cancel-btn', 'Cancel');
    cancel_btn.type = 'button';
    cancel_btn.onclick = () => {
        modal.remove();
        on_cancel();
    };

    const submit_btn = create_element('button', 'submit-btn', 'Create Post');
    submit_btn.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const title = title_input.value.trim();
        const category = category_input.value.trim();
        const content = content_textarea.value.trim();
        
        if (title && category && content) {
            on_submit({ title, category, content });
            modal.remove();
        }
    });

    button_group.append(cancel_btn, submit_btn);
    form.append(title_input, category_input, content_textarea, button_group);
    modal_content.append(modal_header, form);
    modal.append(modal_content);
    document.body.append(modal);
}

export function show_notification(message, type = 'info') {
    const notification = create_element('div', `notification notification-${type}`);
    notification.textContent = message;
    
    document.body.append(notification);
    
    setTimeout(() => {
        notification.classList.add('show');
    }, 10);
    
    setTimeout(() => {
        notification.classList.remove('show');
        setTimeout(() => notification.remove(), 300);
    }, 3000);
}

export function show_loading() {
    const loading = create_element('div', 'loading-overlay');
    const spinner = create_element('div', 'spinner');
    loading.append(spinner);
    loading.id = 'loading-overlay';
    document.body.append(loading);
}

export function hide_loading() {
    const loading = document.getElementById('loading-overlay');
    if (loading) {
        loading.remove();
    }
}