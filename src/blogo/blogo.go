package main

import (
	"encoding/json"

	"net/http"

	"strconv"

	"fmt"

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

	node_articles := api.Root.
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

	node_articles.
		Node("{article_id}").
		Method("GET", func(c *golax.Context) {
			article_id, err := strconv.Atoi(c.Parameters["article_id"])

			if err != nil {
				c.Response.WriteHeader(http.StatusBadRequest)
				fmt.Println("Error parsing URL parameter `article_id`", err)
				return
			}

			if article_id >= len(articles) || article_id < 0 {
				c.Response.WriteHeader(http.StatusNotFound)
				return
			}

			a := articles[article_id]

			json.NewEncoder(c.Response).Encode(a)

		})

	api.Serve()
}
