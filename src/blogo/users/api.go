package users

import (
	"encoding/json"

	"net/http"

	"fmt"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func Build(parent *golax.Node, users_dao *kip.Dao) {

	users_node := parent.Node("users")
	users_node.Method("GET", func(c *golax.Context) {

		result := []interface{}{}

		err := users_dao.
			Find(bson.M{}).
			Limit(20).
			ForEach(func(item *kip.Item) {
				user := item.Value.(*User)
				result = append(result, map[string]interface{}{
					"nick":             user.Nick,
					"create_timestamp": user.CreateTimestamp,
				})
			})

		if nil != err {
			fmt.Println("ERROR:", err)
			c.Error(http.StatusInternalServerError, "Could not read from persistence layer")
			return
		}

		json.NewEncoder(c.Response).Encode(result)
	})

	me_node := users_node.Node("me")
	me_node.Method("GET", func(c *golax.Context) {
		user := GetUser(c)

		if nil == user {
			c.Error(http.StatusForbidden, "You are not logged in")
			return
		}

		json.NewEncoder(c.Response).Encode(user)
	})

}
