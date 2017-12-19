package sessions

import (
	"net/http"
	"strconv"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func CreateSession(sessions_dao *kip.Dao, c *golax.Context) {
	last_guest++

	session_item := sessions_dao.Create()

	session := session_item.Value.(*Session)
	session.UserId = "Invitado " + strconv.Itoa(last_guest)

	if err := session_item.Save(); nil != err {
		c.Error(http.StatusInternalServerError, "Could not save session data")
		return
	}

	c.Set("session", session)

	c.Response.Header().Set("Set-Cookie", COOKIE_NAME+"="+session.Cookie)
}
