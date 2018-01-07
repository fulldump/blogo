package home

import (
	"blogo/articles"

	"html/template"

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
				font-family: sans-serif;
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
		</script>
	</head>
	<body>
		<div class="content">

			<h1><img src="data:image/png;base64,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">BloGo</h1>

			{{range .}}
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

		l := []interface{}{}

		articles_dao.Find(nil).ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			l = append(l, map[string]string{
				"id":      a.Id,
				"title":   a.Title,
				"content": a.Content,
			})
		})

		t.Execute(c.Response, l)

	})

}
