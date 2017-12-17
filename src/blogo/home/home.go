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
		</script>
	</head>
	<body>
		<div class="content">

			<h1>Hola mundo</h1>

			{{range .}}
			<div id="{{.id}}">
				<h2>{{.title}} <button class="button-remove" onclick="removeArticle('{{.id}}')">Borrar</button> </h2>
				<p>{{.content}}</p>
			</div>
			{{end}}

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
