function removeArticle(id) {
    var xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/v0/articles/'+id, true);
    xhr.onload = function() {
        document.getElementById(id).style.display = 'none';
    };
    xhr.send(null);
}

function createArticle() {

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/v0/articles', true);
    xhr.onload = function() {
        window.location.href = '/';
    };

    var title = document.getElementById("create-form-title");
    var content = document.getElementById("create-form-content");

    var payload = {
        "title": title.value,
        "content": content.value,
    };

    xhr.send(JSON.stringify(payload));
}

function formLoginShowToggle(e) {
    var login_panel = document.getElementById('login-panel');

    if (login_panel.style.display === '') {
        login_panel.style.display = 'block';
    } else {
        login_panel.style.display = '';
    }
}

function login(e) {
    e.preventDefault();

    var email = document.getElementById("login-email");
    var password = document.getElementById("login-password");

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/login/email', true);
    xhr.onload = function() {
        window.location.reload();
    };

    var payload = {
        email: email.value,
        password: password.value,
    };

    xhr.send(JSON.stringify(payload));
}

function logout(e) {
    e.preventDefault();

    var xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/v0/sessions/current', true);
    xhr.onload = function() {
        window.location.reload();
    };

    xhr.send();
}

function anchorizeArticles() {
    document.querySelectorAll('section.article p').forEach(function (item) {
        item.innerHTML = anchorme.js(item.innerHTML);
    });
}