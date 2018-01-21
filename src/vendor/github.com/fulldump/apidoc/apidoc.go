package apidoc

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fulldump/golax"
)

type Apidoc struct {
	Title    string
	Subtitle string
}

func Build(a *golax.Api, parent *golax.Node) *Apidoc {

	apidoc := &Apidoc{
		Title:    "Apidoc",
		Subtitle: "Api documentation",
	}

	doc := parent.
		Node("doc").
		Doc(golax.Doc{Ommit: true, Description: `
			This API sub tree is dedicated to self-documentation. It contains documentation
			in several formats: markdown, html and swagger at this moment.

			**Permissions**:
			Anyone can read this.
		`})

	doc.Method("GET", func(c *golax.Context) {
		if !strings.HasSuffix(c.Request.URL.Path, "/") {
			http.Redirect(c.Response, c.Request, c.Request.URL.Path+"/", 302)
			return
		}
		readfile(c, "index.html")
	}, golax.Doc{
		Description: `
			Serve the index file.
		`,
	})

	static := doc.
		Node("static").
		Doc(golax.Doc{Description: `
			Dir for static files
		`})

	static.
		Node("{{*}}").
		Method("GET", func(c *golax.Context) {
			readfile(c, c.Parameter)
		})

	doc.
		Node("json").
		Doc(golax.Doc{Description: `
			Generate documentation in json format.
		`}).
		Method("GET", func(c *golax.Context) {

			c.Response.Header().Set("Content-Type", "application/json")

			j := NewNodeJson(a)
			RunJsonDoc(j)
			j.JsonDoc.Title = apidoc.Title
			j.JsonDoc.Subtitle = apidoc.Subtitle
			json.NewEncoder(c.Response).Encode(j.JsonDoc)
		})

	doc.
		Node("md").
		Doc(golax.Doc{Description: `
			Documentation in markdown format.
		`}).
		Method("GET", func(c *golax.Context) {
			PrintApiMd(NodePrint{
				Api:             a,
				Node:            a.Root,
				Context:         c,
				Path:            a.Prefix,
				AllInterceptors: map[*golax.Interceptor]*golax.Interceptor{},
			})
		})

	doc.
		Node("html").
		Doc(golax.Doc{Description: `
			Documentation in html format.
		`}).
		Method("GET", func(c *golax.Context) {
			c.Error(501, "Unimplemented")
		})

	swagger := doc.
		Node("swagger").
		Doc(golax.Doc{Description: `
			Documentation in swagger format.
		`})
	swagger.Method("GET", func(c *golax.Context) {
		c.Error(501, "Unimplemented")
	})

	return apidoc
}
