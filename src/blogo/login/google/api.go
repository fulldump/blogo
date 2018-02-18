package google

import (
	"fmt"
	"net/http"

	"googleapi"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"

	"blogo/sessions"
	"blogo/users"
)

func Build(parent *golax.Node, dao_users, dao_sessions *kip.Dao, g *googleapi.GoogleApi) {

	google_node := parent.Node("google")

	callback_node := google_node.Node("callback")
	callback_node.Method("GET", func(c *golax.Context) {

		code := c.Request.URL.Query().Get("code")
		state := c.Request.URL.Query().Get("state")

		access, err := g.GetAccessTokenWithHost(code, c.Request.Host)
		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusBadGateway, "Google Auth GetAccessToken communication error")
			return
		}

		user_info, err := access.GetUserInfo()
		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusBadGateway, "Google Auth GetUserInfo communication error")
			return
		}

		user_item, err := dao_users.FindOne(bson.M{"login_google.id": user_info.Id})
		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Could not read from persistence layer")
			return
		}

		if nil == user_item {
			user_item = dao_users.Create()
			user := user_item.Value.(*users.User)
			user.Nick = user_info.GivenName
			user.LoginGoogle = user_info

			err := user_item.Save()
			if nil != err {
				fmt.Println("ERROR:", err)
				c.Error(http.StatusInternalServerError, "Could not write to persistence layer")
				return
			}
		}

		sessions.CreateSession(dao_sessions, c)

		user := user_item.Value.(*users.User)
		if err := sessions.SetSessionUserId(c, user.Id); nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Could not set user session")
			return
		}

		c.Response.Header().Set("Location", state)
		c.Response.WriteHeader(http.StatusFound)

	})

}
