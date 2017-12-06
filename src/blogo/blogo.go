package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fulldump/golax"
)

type Article struct {
	Title   string
	Content string
}

var articles = map[string]Article{
	"uno": {
		Title:   "Uno",
		Content: "1111111",
	},
	"dos": {
		Title:   "Dos",
		Content: "2222222",
	},
}

func main() {

	api := golax.NewApi()

	node_articles := api.Root.
		Node("articles").
		Method("GET", func(c *golax.Context) {

			l := []Article{}

			for _, a := range articles {
				l = append(l, a)
			}

			json.NewEncoder(c.Response).Encode(l)

		}).
		Method("POST", func(c *golax.Context) {

			a := Article{}

			json.NewDecoder(c.Request.Body).Decode(&a)

			u := strings.ToLower(a.Title)

			articles[u] = a

			c.Response.WriteHeader(http.StatusCreated)

		})

	node_articles.
		Node("{article_id}").
		Method("GET", func(c *golax.Context) {
			article_id := c.Parameters["article_id"]

			a, found := articles[article_id]

			if !found {
				c.Response.WriteHeader(http.StatusNotFound)
				return
			}

			json.NewEncoder(c.Response).Encode(a)

		}).
		Method("DELETE", func(c *golax.Context) {
			article_id := c.Parameters["article_id"]

			_, found := articles[article_id]

			if !found {
				c.Response.WriteHeader(http.StatusNotFound)
				return
			}

			delete(articles, article_id)

			c.Response.WriteHeader(http.StatusNoContent)
		})

	api.Serve()
}
