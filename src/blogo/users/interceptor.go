package users

import (
	"blogo/sessions"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

// If a user is logged in, we get the user and put it in context
func NewUserInterceptor(dao *kip.Dao) *golax.Interceptor {
	return &golax.Interceptor{
		Before: func(c *golax.Context) {

			s := sessions.GetSession(c)
			if nil == s {
				return
			}

			user_id := s.Value.(*sessions.Session).UserId

			user_item, err := dao.FindOne(bson.M{"_id": user_id})
			if nil != err {
				// TODO: log: we can not read user from db
				return
			}

			if nil == user_item {
				// TODO: log: user not found in db
				return
			}

			c.Set("user_item", user_item)
		},
	}
}

func GetUser(c *golax.Context) *User {
	user_item, ok := c.Get("user_item")
	if !ok {
		return nil
	}

	return user_item.(*kip.Item).Value.(*User)
}
