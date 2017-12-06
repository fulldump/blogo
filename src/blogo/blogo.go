package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/fulldump/goconfig"
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	Id      string `bson:"_id" json:"id"`
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
}

type Config struct {
	MongoUri string `usage:"MongoDB URI service"`
	HttpAddr string `usage:"TCP port to listen"`
}

func main() {

	// Get configuration
	c := &Config{
		MongoUri: "mongodb://localhost:27017/blogo",
	}

	goconfig.Read(c)

	// Define Dao Articles
	kip.Define(&kip.Collection{
		Name: "articles",
		OnCreate: func() interface{} {
			return &Article{
				Id: bson.NewObjectId().Hex(),
			}
		},
	})

	db, db_err := kip.NewDatabase(c.MongoUri)
	if nil != db_err {
		panic(db_err)
	}

	articles_dao := kip.NewDao("articles", db)

	// Buid API
	api := golax.NewApi()

	node_articles := api.Root.
		Node("articles").
		Method("GET", func(c *golax.Context) {

			l := []interface{}{}

			articles_dao.Find(nil).ForEach(func(item *kip.Item) {
				l = append(l, item.Value)
			})

			json.NewEncoder(c.Response).Encode(l)

		}).
		Method("POST", func(c *golax.Context) {

			item := articles_dao.Create()

			json.NewDecoder(c.Request.Body).Decode(&item.Value)

			if err := item.Save(); nil != err {
				fmt.Println(err)
				c.Response.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.Response.WriteHeader(http.StatusCreated)

		})

	node_articles.
		Node("{article_id}").
		Method("GET", func(c *golax.Context) {

			article_id := c.Parameters["article_id"]

			item, err := articles_dao.FindById(article_id)
			if nil != err {
				fmt.Println(err)
				c.Response.WriteHeader(http.StatusInternalServerError)
				return
			}

			if item == nil {
				c.Response.WriteHeader(http.StatusNotFound)
				return
			}

			json.NewEncoder(c.Response).Encode(item.Value)

		}).
		Method("DELETE", func(c *golax.Context) {

			article_id := c.Parameters["article_id"]

			item, err := articles_dao.FindById(article_id)
			if nil != err {
				fmt.Println(err)
				c.Response.WriteHeader(http.StatusInternalServerError)
				return
			}

			if item == nil {
				c.Response.WriteHeader(http.StatusNotFound)
				return
			}

			if err := item.Delete(); nil != err {
				fmt.Println(err)
				c.Response.WriteHeader(http.StatusInternalServerError)
				return
			}

			c.Response.WriteHeader(http.StatusNoContent)
		})

	api.Serve()
}
