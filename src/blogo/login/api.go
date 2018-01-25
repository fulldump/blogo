package login

import (
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"

	"googleapi"

	"blogo/login/email"
	"blogo/login/google"
)

func Build(parent *golax.Node, dao_users, dao_sessions *kip.Dao, g *googleapi.GoogleApi) {

	login_node := parent.Node("login")

	email.Build(login_node, dao_users, dao_sessions)
	google.Build(login_node, dao_users, dao_sessions, g)

}
