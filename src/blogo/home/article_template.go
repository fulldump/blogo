package home

var article_template = `
{{define "content"}}
			<h2>{{.article.Title}}</h2>
			<p>{{.article.Content}}</p>
{{end}}
`
