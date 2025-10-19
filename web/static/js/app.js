import { render_login_form, render_home_page, render_posts } from './ui.js';
import { check_auth_status, login_user, get_posts } from './api.js';

async function main() {
    const is_logged_in = await check_auth_status();

    if (is_logged_in) {
        show_home_page();
    } else {
        show_login_page();
    }
}

async function show_home_page() {
    render_home_page();
    const posts = await get_posts();
    render_posts(posts);
}

function show_login_page() {
    render_login_form(async (form_data) => {
        const success = await login_user(form_data);
        if (success) {
            show_home_page();
        } else {
            alert('Login failed! Please try again.');
        }
    }, () => {
        alert('Registration form coming soon!');
    });
}

main();
