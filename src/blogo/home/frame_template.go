package home

var frame_template = `
<!DOCTYPE html>
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
					{{if .user.LoginGoogle }}
					<img src="{{.user.LoginGoogle.Picture}}" style="height: 24px; vertical-align: text-bottom;">
					{{end}}
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
				<a href="{{ .google_oauth_link }}">Entrar con Google</a>
				<script>
					document.getElementById('form-login').addEventListener('submit', login, true);
				</script>
				{{end}}
			</div>

			<a href="/"><h1>BloGo</h1></a>

			{{block "content" .}}
				THIS IS THE CONTENT
			{{end}}
		<div>
	</body>
</html>
`
