package sessions

import (
	"time"

	"github.com/fulldump/goaudit"
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

const COOKIE_NAME = "blogo"

func NewSessionInterceptor(dao *kip.Dao) *golax.Interceptor {
	return &golax.Interceptor{
		Before: func(c *golax.Context) {

			audit := goaudit.GetAudit(c)

			cookie, err := c.Request.Cookie(COOKIE_NAME)

			if nil != err {
				audit.Log.Info("No cookie")
				CreateSession(dao, c)
				return
			}

			cookie_hash := hash(cookie.Value)
			session_item, err := dao.FindOne(bson.M{"cookie": cookie_hash})
			if nil != err {
				audit.Log.Error("Can not read from db", err)
				CreateSession(dao, c)
				return
			}

			if nil == session_item {
				audit.Log.Info("Session not found")
				CreateSession(dao, c)
				return
			}

			session := session_item.Value.(*Session)

			now := time.Now().Unix()
			if now > session.ExpireTimestamp {
				audit.Log.Info("Expired session")
				CreateSession(dao, c)
				return
			}

			audit.SessionId = session.Id

			c.Set("session", session_item)
		},
		After: func(c *golax.Context) {

			audit := goaudit.GetAudit(c)

			value, ok := c.Get("session")
			if !ok {
				audit.Log.Fatal("Session does not exist!")
				return
			}

			session_item := value.(*kip.Item)
			session_item.Patch(&kip.Patch{
				Key:       "data.navigation",
				Operation: "add_to_set",
				Value: map[string]interface{}{
					"method":    c.Request.Method,
					"path":      c.Request.URL.Path,
					"timestamp": time.Now().Unix(),
				},
			})
			err := session_item.Save()
			if nil != err {
				audit.Log.Error(err)
			}
		},
	}
}

func CreateSession(dao *kip.Dao, c *golax.Context) (session_item *kip.Item) {

	cookie, session_item := NewSession(dao)

	audit := goaudit.GetAudit(c)

	if err := session_item.Save(); err != nil {
		audit.Log.Error(err)
		return
	}

	session := session_item.Value.(*Session)

	audit.SessionId = session.Id
	audit.Log.Info("New session")

	c.Response.Header().Set("Set-Cookie", COOKIE_NAME+"="+cookie+"; Path=/")

	c.Set("session", session_item)

	return
}

func GetSession(c *golax.Context) *kip.Item {

	value, ok := c.Get("session")
	if !ok {
		// todo: log, session does not exist
		return nil
	}

	return value.(*kip.Item)
}

func SetSessionUserId(c *golax.Context, user_id string) error {
	s := GetSession(c)
	s.Patch(&kip.Patch{
		Operation: "set",
		Key:       "user_id",
		Value:     user_id,
	})
	return s.Save()
}
