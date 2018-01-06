package sessions

import (
	"encoding/json"

	"net/http"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func Build(parent *golax.Node, sessions_dao *kip.Dao) {

	sessions_node := parent.Node("sessions")
	sessions_node.Method("GET", func(c *golax.Context) {

		result := []interface{}{}

		sessions_dao.
			Find(bson.M{}).
			Sort("-create_timestamp").
			Limit(20).
			ForEach(func(item *kip.Item) {
				result = append(result, item.Value)
			})

		json.NewEncoder(c.Response).Encode(result)

	})

	current_node := sessions_node.Node("current")
	current_node.Method("GET", func(c *golax.Context) {

		value, ok := c.Get("session")
		if !ok {
			// todo: log, session does not exist
			c.Error(http.StatusNotFound, "Session does not exist")
			return
		}

		session_item := value.(*kip.Item)

		json.NewEncoder(c.Response).Encode(session_item.Value)

	})

}
