package email

import (
	"blogo/sessions"
	"blogo/users"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func Build(parent *golax.Node, dao_users *kip.Dao) {

	email_node := parent.Node("email")
	email_node.Method("POST", func(c *golax.Context) {

		body := struct {
			Email    *string `json:"email"`
			Password *string `json:"password"`
		}{}

		err := json.NewDecoder(c.Request.Body).Decode(&body)
		if nil != err {
			fmt.Println("INFO:", err)
			c.Error(http.StatusBadRequest, "Malformed JSON")
			return
		}

		if nil == body.Email {
			c.Error(http.StatusBadRequest, "Field `email` is mandatory")
			return
		}

		if nil == body.Password {
			c.Error(http.StatusBadRequest, "Field `password` is mandatory")
			return
		}

		user_item, err := dao_users.FindOne(bson.M{"login_email.email": body.Email})
		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Could not read from persistence layer")
			return
		}

		if nil == user_item {
			c.Error(http.StatusForbidden, "Bad credentials")
			return
		}

		user := user_item.Value.(*users.User)
		if !user.LoginEmail.Check(*body.Password) {
			c.Error(http.StatusForbidden, "Bad credentials")
			return
		}

		if err := sessions.SetSessionUserId(c, user.Id); nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Could not set user session")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)

	})

}
