package sessions

import (
	"encoding/json"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

var last_guest = 1

func Build(parent *golax.Node, sessions_dao *kip.Dao) {

	sessions := parent.Node("sessions")

	sessions.Method("GET", func(c *golax.Context) {

		session := GetSession(c)

		l := []Session{}

		sessions_dao.Find(bson.M{"user_id": session.UserId}).All(&l)

		json.NewEncoder(c.Response).Encode(l)

	})

}
