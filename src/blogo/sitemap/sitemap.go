package sitemap

import (
	"blogo/articles"
	"time"

	"fmt"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(parent *golax.Node, articles_dao *kip.Dao) {

	parent.Node("sitemap.xml").
		Method("GET", func(c *golax.Context) {

			c.Response.Header().Set("Content-Type", "text/xml; charset=UTF-8")

			c.Response.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" 
  xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" 
  xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
`))

			fmt.Println(c.Request.Header, "proto:", c.Request.Proto)

			proto := c.Request.Header.Get("X-Forwarded-Proto")

			host := c.Request.Header.Get("X-Forwarded-Host")

			url_prefix := proto + "://" + host

			articles_dao.Find(nil).ForEach(func(article_item *kip.Item) {

				article := article_item.Value.(*articles.Article)

				loc := url_prefix + "/@" + article.User.Nick + "/" + article.TitleUrl
				lastmod := time.Unix(0, article.CreateTimestamp).Format(time.RFC3339)

				c.Response.Write([]byte(`<url>
	<loc>` + loc + `</loc>
	<lastmod>` + lastmod + `</lastmod>
</url>
`))
			})

			c.Response.Write([]byte(`</urlset>`))

		})

}
