// export function hello_from_ui() {
//     console.log("UI ..........aldlaslkdhasld");
// }


function create_element(tag, className, text) {
    const el = document.createElement(tag)
    if (className) el.className = className
    if (text) el.textContent = text
    return el
}

export function render_login_form(on_login , on_switch_to_register){
    document.body.innerHTML = ""

    const c = create_element('div' , 'form-container')
    const title = create_element('h1', '' , 'login')

    const form = create_element('form')
    const username_inpt = create_element('input')
    username_inpt.placeholder = 'Username or Email'
    username_inpt.name = 'username'

    const passwd_inpt = create_element('input')
    passwd_inpt.type = 'password'
    passwd_inpt.placeholder = 'password'
    passwd_inpt.name = 'password'

    const sub_button = create_element('button', '', 'Login');
    sub_button.type = 'submit';

    form.addEventListener('submit', (e) => {
        e.preventDefault(); 
        const form_data = new FormData(form)
        on_login(form_data)
    })
    const togl_txt = create_element('p' , 'toggle-form' , "Don't have an account? Register")
    togl_txt.onclick = on_switch_to_register

    form.append(username_inpt,passwd_inpt,sub_button)
    c.append(title,form,togl_txt)

document.body.append(c)
}
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

export function render_home_page() {
    document.body.innerHTML = '';
    document.body.className = 'home-layout';

    const main_container = create_element('div', 'main-container');
    const page_title = create_element('h1', 'page-title', 'Forum Posts');
    
    const posts_container = create_element('div', 'posts-list');
    posts_container.id = 'posts-container';

    main_container.append(page_title, posts_container);
    document.body.append(main_container);
}
