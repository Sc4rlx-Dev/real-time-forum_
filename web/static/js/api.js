// export function hello_from_api() {
//     console.log("Api : safasfasfasf")
// }


export async function check_auth_status() {
    return document.cookie.includes('session_token=')
}

export async function login_user(form_data) {
    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            body: form_data,
        })
        return response.ok
    } catch (error) { console.error('Login failed:', error);
        return false
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