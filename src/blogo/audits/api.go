package audits

import (
	"blogo/users"
	"encoding/json"
	"net/http"

	"github.com/fulldump/goaudit"
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2/bson"
)

func Build(parent *golax.Node, audits_dao *kip.Dao) {

	audits := parent.Node("audits")

	audits.Interceptor(&golax.Interceptor{
		Before: func(c *golax.Context) {

			user := users.GetUser(c)

			if nil == user {
				c.Error(http.StatusUnauthorized, "Invalid user")
				return
			}

			if !user.Scopes.Admin {
				c.Error(http.StatusForbidden, "Your user is not authorized")
				return
			}

		},
	})

	audits.Method("GET", func(c *golax.Context) {

		goaudit.GetAudit(c).Abort()

		limit := 20 // TODO: overwrite with get param `limit`

		filter := bson.M{}

		filters := c.Request.URL.Query()["filter"]
		filtersLen := len(filters)
		if filtersLen == 1 {
			if err := json.Unmarshal([]byte(filters[0]), &filter); nil != err {
				c.Error(http.StatusBadRequest, "Could not parse 'filter' parameter from url. Expected JSON object.")
				return
			}
		} else if filtersLen > 1 {
			c.Error(http.StatusBadRequest, "More than one 'filter' parameter received")
			return
		}

		encoder := json.NewEncoder(c.Response)
		audits_dao.Find(filter).Limit(limit).Sort("-entry_timestamp").ForEach(func(item *kip.Item) {
			encoder.Encode(item.Value)
		})

	})

	audit := audits.Node("{audit_id}")

	audit.Method("GET", func(c *golax.Context) {

		audit := goaudit.GetAudit(c)

		audit_id := c.Parameters["audit_id"]
		audit_item, err := audits_dao.FindById(audit_id)
		if nil != err {
			audit.Log.Error(err)
			return
		}

		if nil == audit_item {
			c.Error(http.StatusNotFound, "Audit not found")
			return
		}

		encoder := json.NewEncoder(c.Response)
		encoder.SetIndent("", "    ")
		encoder.Encode(audit_item.Value)
	})

}
