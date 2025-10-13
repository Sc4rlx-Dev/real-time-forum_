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

export function render_home_page(){
    document.body.innerHTML = ''
    const tittle = create_element('h1' , '' , 'Mer7ba biiiiiiik !!')
    document.body.append(tittle)
}