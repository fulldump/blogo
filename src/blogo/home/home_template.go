package home

var home_template = `
{{define "content"}}

			{{range .articles}}
			<div id="{{.Id}}">
				<h2><a href="/a/{{.Id}}">{{.Title}}</a> <button class="button-remove" onclick="removeArticle('{{.Id}}')">Borrar</button> </h2>
				<p>{{.Content}}</p>
			</div>
			{{end}}

			{{if .user}}
			<div class="create-form">
				<input type="text" id="create-form-title" placeholder="Título">
				<textarea type="text" id="create-form-content" placeholder="Contenido..."></textarea>

				<div style="text-align: center;">
					<button id="create-form-button" class="button-create" onclick="createArticle()">Crear</button>
				</div>
			</div>
			{{else}}
			<div style="text-align: center;">Identifícate para crear artículos</div>
			{{end}}

		<div>
{{end}}
`
