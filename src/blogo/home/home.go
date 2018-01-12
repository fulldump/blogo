package home

import (
	"html/template"

	"blogo/users"

	"blogo/articles"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(parent *golax.Node, articles_dao *kip.Dao) {

	t, _ := template.New("home").Parse(html)

	parent.Method("GET", func(c *golax.Context) {

		user := users.GetUser(c)

		articles_list := []interface{}{}

		articles_dao.Find(nil).ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			articles_list = append(articles_list, map[string]string{
				"id":      a.Id,
				"title":   a.Title,
				"content": a.Content,
			})
		})

		t.Execute(c.Response, map[string]interface{}{
			"user":     user,
			"articles": articles_list,
		})

	})

}
