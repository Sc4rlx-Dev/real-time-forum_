// console.log("tesstttt")

import { render_login_form, render_home_page } from './ui.js';
import { check_auth_status, login_user } from './api.js';

async function main(){
    const is_logged_in = await check_auth_status()
    if(is_logged_in){
        show_home_page()
    }else{show_login_page()}
}

function show_home_page(){
    render_home_page()
}

function show_login_page(){
    render_login_form(async (form_data) => {
        const sucss = await login_user(form_data)
        if(sucss){
            show_home_page
        }else{alert('wa  l3adeeew')}
    }, () => {
        alert('registration soon..')
    })
}
main()