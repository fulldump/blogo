package login

import (
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"

	"googleapi"

	"blogo/login/email"
	"blogo/login/google"
)

func Build(parent *golax.Node, dao_users *kip.Dao, g *googleapi.GoogleApi) {

	login_node := parent.Node("login")

	email.Build(login_node, dao_users)
	google.Build(login_node, dao_users, g)

}
