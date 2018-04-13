package background

import (
	"blogo/articles"
	"fmt"

	"github.com/fulldump/kip"
)

func TitleUrlize(articles_dao *kip.Dao) {

	articles_dao.Find(nil).ForEach(func(article_item *kip.Item) {

		article := article_item.Value.(*articles.Article)

		article_item.
			Patch(&kip.Patch{
				Operation: "set",
				Key:       "title_url",
				Value:     articles.UrlizeTitle(article.TitleUrl),
			})

		err := article_item.Save()
		if nil != err {
			fmt.Println("Error saving article:", err)
		}
	})

}
