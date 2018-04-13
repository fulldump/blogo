package home

import (
	"fmt"
	"googleapi"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"

	"blogo/articles"
	"blogo/httputils"
	"blogo/statics"
	"blogo/users"
)

func p(name string, codes ...string) (t *template.Template, err error) {

	t = template.New(name)

	for _, code := range codes {
		t, err = t.Parse(code)
		// TODO: return on error
	}

	return
}

func Build(parent *golax.Node, articles_dao, users_dao *kip.Dao, g *googleapi.GoogleApi, google_analytics string) {

	template_html := string(statics.Bytes["template.html"])
	index_html := string(statics.Bytes["index.html"])
	article_html := string(statics.Bytes["article.html"])

	t_home, _ := p("home", template_html, index_html)
	t_article, _ := p("article", template_html, article_html)

	parent.Method("GET", func(c *golax.Context) {

		user := users.GetUser(c)

		articles_list := []interface{}{}

		articles_dao.Find(nil).Sort("-create_timestamp").ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			articles_list = append(articles_list, formatArticleData(a))
		})

		err := t_home.Execute(c.Response, map[string]interface{}{
			"user":              user,
			"articles":          articles_list,
			"google_oauth_link": g.CreateLinkWithHost(c.Request.URL.Path, httputils.GetHost(c.Request)),
			"google_analytics":  google_analytics,
		})

		if nil != err {
			fmt.Println("ERROR:", err)
			return
		}

	})

	user_regex := "^@[^@]+"

	user_arroba := parent.Node("(" + user_regex + ")")
	user_arroba.Interceptor(&golax.Interceptor{
		Before: func(c *golax.Context) {

			user_nick := strings.TrimLeft(c.Parameters[user_regex], "@")

			q := bson.M{
				"nick": user_nick,
			}

			user_item, err := users_dao.FindOne(q)
			if nil != err {
				c.Error(http.StatusInternalServerError, "Unexpected error reading from persistence layer")
				return
			}

			if nil == user_item {
				// TODO: pretify this with an html page instead of a fucking unfriendly json
				c.Error(http.StatusNotFound, "User not found!!!")
				return
			}

			c.Set("blog_user", user_item.Value)

		},
	})

	user_arroba.Method("GET", func(c *golax.Context) {

		user := GetBlogUser(c)

		articles_list := []interface{}{}

		q := bson.M{
			"user.nick": user.Nick,
		}

		articles_dao.Find(q).Sort("-create_timestamp").ForEach(func(item *kip.Item) {
			a := item.Value.(*articles.Article)

			articles_list = append(articles_list, formatArticleData(a))
		})

		err := t_home.Execute(c.Response, map[string]interface{}{
			"user":              users.GetUser(c),
			"articles":          articles_list,
			"google_oauth_link": g.CreateLinkWithHost(c.Request.URL.Path, httputils.GetHost(c.Request)),
			"google_analytics":  google_analytics,
		})

		if nil != err {
			fmt.Println("ERROR:", err)
			return
		}

	})

	user_arroba.Node("{{title_url}}").Method("GET", func(c *golax.Context) {

		user := GetBlogUser(c)
		title_url := c.Parameters["title_url"]

		q := bson.M{
			"user.nick": user.Nick,
			"title_url": title_url,
		}

		article_item, err := articles_dao.FindOne(q)
		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Unexpected error reading from persitence layer.")
			return
		}

		if nil == article_item {
			c.Error(http.StatusNotFound, "Article not found") // TODO: maybe a pretty page is better :D
			return
		}

		article := article_item.Value.(*articles.Article)

		err = t_article.Execute(c.Response, map[string]interface{}{
			"user":              users.GetUser(c),
			"article":           formatArticleData(article),
			"google_oauth_link": g.CreateLinkWithHost(c.Request.URL.Path, httputils.GetHost(c.Request)),
			"google_analytics":  google_analytics,
		})

	})

}

func GetBlogUser(c *golax.Context) *users.User {
	item, exist := c.Get("blog_user")
	if exist {
		return item.(*users.User)
	}

	return nil
}

func formatArticleData(a *articles.Article) interface{} {
	return map[string]interface{}{
		"Id":              a.Id,
		"User":            a.User,
		"Title":           a.Title,
		"TitleUrl":        a.TitleUrl,
		"Content":         a.Content,
		"CreateTimestamp": a.CreateTimestamp,
		"CreateDate":      formatTime(time.Unix(0, a.CreateTimestamp)),
	}
}

var months = map[time.Month]string{
	time.January:   "Enero",
	time.February:  "Febrero",
	time.March:     "Marzo",
	time.April:     "Abril",
	time.May:       "Mayo",
	time.June:      "Junio",
	time.July:      "Julio",
	time.August:    "Agosto",
	time.September: "Septiembre",
	time.October:   "Octubre",
	time.November:  "Noviembre",
	time.December:  "Diciembre",
}

func formatTime(t time.Time) (f string) {

	t = t.UTC()

	f = strconv.Itoa(t.Day()) + " " + months[t.Month()]

	now := time.Now()
	if now.Year() != t.Year() {
		f += ", " + strconv.Itoa(t.Year())
	}

	return
}
