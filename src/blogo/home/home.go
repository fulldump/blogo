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
		</style>
	</head>
	<body>
		<div class="content">

			<h1>Hola mundo</h1>

			{{range .}}
				<h2>{{.title}}</h2>
				<p>{{.content}}</p>
			{{end}}

		<div>
	</body>
</html>`)

	parent.Method("GET", func(c *golax.Context) {

		l := []interface{}{}

		articles_dao.Find(nil).ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			l = append(l, map[string]string{
				"title":   a.Title,
				"content": a.Content,
			})
		})

		t.Execute(c.Response, l)

	})

}
