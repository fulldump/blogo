package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"

	"blogo/users"
)

func Build(parent *golax.Node, articles_dao *kip.Dao) {

	articles := parent.Node("articles")
	article := articles.Node("{article_id}")

	articles.Method("GET", func(c *golax.Context) {

		l := []interface{}{}

		articles_dao.Find(nil).ForEach(func(item *kip.Item) {
			l = append(l, item.Value)
		})

		json.NewEncoder(c.Response).Encode(l)

	})

	articles.Method("POST", func(c *golax.Context) {

		user := users.GetUser(c)

		if user == nil {
			c.Error(http.StatusUnauthorized, "You should be logged in")
			return
		}

		if user.Scopes.Banned {
			c.Error(http.StatusForbidden, "You are banned! Banned users can not create articles.")
			return
		}

		item := articles_dao.Create()

		json.NewDecoder(c.Request.Body).Decode(&item.Value)

		article := item.Value.(*Article)
		article.TitleUrl = UrlizeTitle(article.Title)
		article.CreateTimestamp = time.Now().UnixNano()
		article.User = User{
			Id:            user.Id,
			Nick:          user.Nick,
			SyncTimestamp: time.Now().UnixNano(),
		}

		if err := item.Save(); nil != err {
			fmt.Println(err)
			c.Response.WriteHeader(http.StatusInternalServerError)
			return
		}

		c.Response.WriteHeader(http.StatusCreated)

	})

	article.Method("GET", func(c *golax.Context) {

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

	})

	article.Method("DELETE", func(c *golax.Context) {

		user := users.GetUser(c)

		if user == nil {
			c.Error(http.StatusUnauthorized, "You should be logged in")
			return
		}

		if !user.Scopes.Admin {
			c.Error(http.StatusForbidden, "Only admins can remove articles.")
			return
		}

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

}

func UrlizeTitle(s string) string {

	s = strings.Replace(s, " ", "-", -1)

	s = strings.Replace(s, "--", "-", -1)
	s = strings.Replace(s, "--", "-", -1)

	return s
}
