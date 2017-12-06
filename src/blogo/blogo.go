package main

import (
	"encoding/json"

	"net/http"

	"github.com/fulldump/golax"
)

type Article struct {
	Title   string
	Content string
}

var articles = []Article{
	{
		Title:   "Uno",
		Content: "1111111",
	},
	{
		Title:   "Dos",
		Content: "2222222",
	},
}

func main() {

	api := golax.NewApi()

	api.Root.
		Node("articles").
		Method("GET", func(c *golax.Context) {

			json.NewEncoder(c.Response).Encode(articles)

		}).
		Method("POST", func(c *golax.Context) {

			a := Article{}

			json.NewDecoder(c.Request.Body).Decode(&a)

			articles = append(articles, a)

			c.Response.WriteHeader(http.StatusCreated)

		})

	api.Serve()
}
