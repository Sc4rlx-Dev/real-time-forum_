export async function check_auth_status() {
    return document.cookie.includes('session_token=');
}

export async function login_user(form_data) {
    try {
        // Convert FormData to URLSearchParams for application/x-www-form-urlencoded
        const params = new URLSearchParams(form_data);
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: params,
        });
        if (response.ok) {
            const data = await response.json();
            return { success: true, username: data.username };
        }
        return { success: false };
    } catch (error) { 
        console.error('Login failed:', error);
        return { success: false };
    }
}

export async function register_user(form_data) {
    try {
        // Convert FormData to URLSearchParams for application/x-www-form-urlencoded
        const params = new URLSearchParams(form_data);
        const response = await fetch('/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: params,
        });
        return response.ok;
    } catch (error) {
        console.error('Registration failed:', error);
        return false;
    }
}

export async function get_posts() {
    try {
        const response = await fetch('/api/posts');
        if (!response.ok) {
            console.error('Failed to fetch posts:', response.statusText);
            return []; 
        }
        const posts = await response.json();
        return posts || [];
    } catch (error) {
        console.error('Error fetching posts:', error);
        return [];
    }
}

export async function get_users() {
    try {
        const response = await fetch('/api/users');
        if (!response.ok) {
            console.error('Failed to fetch users:', response.statusText);
            return [];
        }
        const users = await response.json();
        return users || [];
    } catch (error) {
        console.error('Error fetching users:', error);
        return [];
    }
}

export async function get_chat_history(username) {
    try {
        const response = await fetch(`/api/messages/${username}`);
        if (!response.ok) {
            console.error('Failed to fetch messages:', response.statusText);
            return [];
        }
        const messages = await response.json();
        return messages || [];
    } catch (error) {
        console.error('Error fetching messages:', error);
        return [];
    }
}

export async function get_conversations() {
    try {
        const response = await fetch('/api/conversations');
        if (!response.ok) {
            console.error('Failed to fetch conversations:', response.statusText);
            return [];
        }
        const conversations = await response.json();
        return conversations || [];
    } catch (error) {
        console.error('Error fetching conversations:', error);
        return [];
    }
}



export function get_cookie(name) {
    const cookies = document.cookie.split(";").map(c => c.trim());
    for (const cookie of cookies) {
        if (cookie.startsWith(name + "=")) {
            return cookie.substring(name.length + 1);
        }
    }
    return null;
}

export async function create_post(post_data) {
    try {
        const response = await fetch('/api/posts/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(post_data),
        });
        return response.ok;
    } catch (error) {
        console.error('Error creating post:', error);
        return false;
    }
}

export async function add_comment(post_id, content) {
    try {
        const response = await fetch('/api/comments/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                post_id: post_id.toString(),
                content: content,
            }),
        });
        return response.ok;
    } catch (error) {
        console.error('Error adding comment:', error);
        return false;
    }
}

export async function logout_user() {
	try {
		await fetch('/api/logout', {
			method: 'POST',
			credentials: 'same-origin'
		});
	} catch (error) {
		console.error('Logout error:', error);
	}
	return true;
}
