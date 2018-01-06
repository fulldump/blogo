package sessions

import (
	"time"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

const COOKIE_NAME = "blogo"

func NewSessionsInterceptor(dao *kip.Dao) *golax.Interceptor {
	return &golax.Interceptor{
		Before: func(c *golax.Context) {

			cookie, err := c.Request.Cookie(COOKIE_NAME)

			if nil != err {
				CreateSession(dao, c)
				return
			}

			session_item, err := GetSession(dao, cookie.Value)
			if nil != err {
				// TODO: log error looking for session
				CreateSession(dao, c)
				return
			}

			if nil == session_item {
				// TODO: log session not found
				CreateSession(dao, c)
				return
			}

			session := session_item.Value.(*Session)

			now := time.Now().Unix()
			if now > session.ExpireTimestamp {
				// Expired session
				// TODO: log expired session
				CreateSession(dao, c)
				return
			}

			c.Set("session", session_item)
		},
		After: func(c *golax.Context) {

			value, ok := c.Get("session")
			if !ok {
				// todo: log, session does not exist
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
			session_item.Save() // TODO: handle error
		},
	}
}

func CreateSession(dao *kip.Dao, c *golax.Context) {
	cookie, session_item := NewSession(dao)

	if err := session_item.Save(); err != nil {
		// TODO: log problem persisting session
		return
	}

	c.Response.Header().Set("Set-Cookie", COOKIE_NAME+"="+cookie)

	c.Set("session", session_item)
}
