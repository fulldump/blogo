package sessions

import (
	"net/http"
	"time"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func NewSessionInterceptor(sessions_dao *kip.Dao) *golax.Interceptor {
	return &golax.Interceptor{
		Before: func(c *golax.Context) {

			cookie, err := c.Request.Cookie(COOKIE_NAME)

			if nil != err {
				CreateSession(sessions_dao, c)
				return
			}

			session_item, err := sessions_dao.FindOne(bson.M{"cookie": cookie.Value})
			if nil != err {
				c.Error(http.StatusInternalServerError, err.Error())
				return
			}

			if nil == session_item {
				CreateSession(sessions_dao, c)
				return
			}

			session := session_item.Value.(*Session)

			now := time.Now().Unix()

			if now > session.ExpirationTimestamp {
				// Expired session
				CreateSession(sessions_dao, c)
				return
			}

			CreateSession(sessions_dao, c)

		},
	}
}

func GetSession(c *golax.Context) *Session {

	value, exists := c.Get("session")

	if !exists {
		panic("Session should exist in context")
	}

	session, ok := value.(*Session)
	if !ok {
		panic("Session in context should be *Session type")
	}

	return session
}
