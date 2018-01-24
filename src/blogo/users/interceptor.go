package users

import (
	"blogo/sessions"

	"github.com/fulldump/goaudit"
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

// If a user is logged in, we get the user and put it in context
func NewUserInterceptor(dao *kip.Dao) *golax.Interceptor {
	return &golax.Interceptor{
		Before: func(c *golax.Context) {

			audit := goaudit.GetAudit(c)

			s := sessions.GetSession(c)
			if nil == s {
				audit.Log.Warning("No session. Session should be always present.")
				return
			}

			user_id := s.Value.(*sessions.Session).UserId

			if "" == user_id {
				return
			}

			user_item, err := dao.FindOne(bson.M{"_id": user_id})
			if nil != err {
				audit.Log.Error("Can not read user from db", err)
				return
			}

			if nil == user_item {
				audit.Log.Infof("User '%s' not found", user_id)
				return
			}

			user := user_item.Value.(*User)

			audit.AuthId = user.Id

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
