package home

import (
	"html/template"

	"blogo/users"

	"blogo/articles"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(parent *golax.Node, articles_dao *kip.Dao) {

	t, _ := template.New("home").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>BloGo</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
			html {
				font-size: 120%;
			}

			.content {
				max-width: 800px;
				margin: auto;
			}

			h1 {
				color: #303060;
				text-align: center;
			}

			h2 {
				color: #3030A0;
			}

			.button-remove {
				background-color: red;
				color:white;
				border: solid #660000 1px;
				border-radius: 3px;
				display: inline-block;
				cursor: pointer;
			}

			.button-create {
				background-color: blue;
				color:white;
				border: solid #000066 1px;
				border-radius: 3px;
				display: inline-block;
				cursor: pointer;
			}

			.create-form {
				border: solid gray 1px;
				border-radius: 4px;
				background-color: #F8F8F8;
				padding: 16px;
			}

			#create-form-title {
				display: block;
				width: 100%;
				font-weight: bold;
				font-size: 150%;
			}

			#create-form-content {
				display: block;
				width: 100%;
			}

			.auth {
				text-align: right;
			}

		</style>

		<script>
			function removeArticle(id) {
				var xhr = new XMLHttpRequest();
				xhr.open('DELETE', '/articles/'+id, true);
				xhr.onload = function() {
					document.getElementById(id).style.display = 'none';
				};
				xhr.send(null);
			}

			function createArticle() {

				var xhr = new XMLHttpRequest();
				xhr.open('POST', '/articles', true);
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

			function login(e) {
				e.preventDefault();

				var email = document.getElementById("login-email");
				var password = document.getElementById("login-password");

				var xhr = new XMLHttpRequest();
				xhr.open('POST', '/login/email', true);
				xhr.onload = function() {
					window.location.href = '/';
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
				xhr.open('DELETE', '/sessions/current', true);
				xhr.onload = function() {
					window.location.href = '/';
				};

				xhr.send();
			}

		</script>
	</head>
	<body>
		<div class="content">

			<div class="auth">
				{{if .user}}
				<form id="form-logout">
					{{.user.Nick}} <button>Salir</button>
				</form>
				<script>
					document.getElementById('form-logout').addEventListener('submit', logout, true);
				</script>
				{{else}}
				<form id="form-login">
					<input type="text" id="login-email" placeholder="tu@email.com">
					<input type="password" id="login-password" placeholder="contraseña">
					<button>Entrar</button>
				</form>
				<script>
					document.getElementById('form-login').addEventListener('submit', login, true);
				</script>
				{{end}}
			</div>

			<h1>BloGo</h1>

			{{range .articles}}
			<div id="{{.id}}">
				<h2>{{.title}} <button class="button-remove" onclick="removeArticle('{{.id}}')">Borrar</button> </h2>
				<p>{{.content}}</p>
			</div>
			{{end}}

			<div class="create-form">
				<input type="text" id="create-form-title" placeholder="Título">
				<textarea type="text" id="create-form-content" placeholder="Contenido..."></textarea>

				<div style="text-align: center;">
					<button id="create-form-button" class="button-create" onclick="createArticle()">Crear</button>
				</div>
			</div>

		<div>
	</body>
</html>`)

	parent.Method("GET", func(c *golax.Context) {

		user := users.GetUser(c)

		articles_list := []interface{}{}

		articles_dao.Find(nil).ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			articles_list = append(articles_list, map[string]string{
				"id":      a.Id,
				"title":   a.Title,
				"content": a.Content,
			})
		})

		t.Execute(c.Response, map[string]interface{}{
			"user":     user,
			"articles": articles_list,
		})

	})

}
