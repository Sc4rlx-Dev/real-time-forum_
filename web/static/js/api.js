// Checks if the user is authenticated by looking for a session token in cookies.
export async function check_auth_status() {
    return document.cookie.includes('session_token=');
}

// Sends the login form data to the backend.
export async function login_user(form_data) {
    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            body: form_data,
        });
        return response.ok;
    } catch (error) { 
        console.error('Login failed:', error);
        return false;
    }
}
// --------------------

// --- NEW FUNCTION ---
// Sends the registration form data to the backend.
export async function register_user(form_data) {
    try {
        const response = await fetch('/api/register', {
            method: 'POST',
            body: form_data,
        });
        return response.ok; // Returns true if registration was successful
    } catch (error) {
        console.error('Registration failed:', error);
        return false;
    }
}
// --------------------

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



export function get_cookie(name) {
  const cookies = document.cookie.split(";").map(c => c.trim());
  for (const cookie of cookies) {
    if (cookie.startsWith(name + "=")) {
      return cookie.substring(name.length + 1);
    }
  }
  return null
}
