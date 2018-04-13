package background

import (
	"time"

	"blogo/articles"
	"blogo/users"
	"fmt"

	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func UsersInArticle(users_dao, articles_dao *kip.Dao) {

	for {

		yesterday := time.Now().Add(-24 * time.Hour).UnixNano()
		now := time.Now().UnixNano()

		q := bson.M{
			"$or": []bson.M{
				{
					"user.sync_timestamp": bson.M{"$exists": false},
				},
				{
					"user.sync_timestamp": bson.M{"$lt": yesterday},
				},
			},
		}

		articles_dao.Find(q).ForEach(func(article_item *kip.Item) {

			article := article_item.Value.(*articles.Article)
			user_id := article.User.Id

			user_item, err := users_dao.FindById(user_id)
			if nil != err {
				return
			}

			if user_item == nil {
				return
			}

			user := user_item.Value.(*users.User)

			article_item.
				Patch(&kip.Patch{
					Operation: "set",
					Key:       "user",
					Value: articles.User{
						Id:            user.Id,
						Nick:          user.Nick,
						SyncTimestamp: now,
					},
				})

			err = article_item.Save()
			if nil != err {
				fmt.Println("Error saving article:", err)
			}
		})

		time.Sleep(1 * time.Minute)
	}

}
