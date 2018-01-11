package api

import (
	"blogo/articles"

	"blogo/home"
	"blogo/sessions"

	"blogo/login"
	"blogo/users"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(articles_dao, sessions_dao, users_dao *kip.Dao) *golax.Api {

	api := golax.NewApi()

	api.Root.Interceptor(sessions.NewSessionInterceptor(sessions_dao))
	api.Root.Interceptor(users.NewUserInterceptor(users_dao))
	api.Root.Interceptor(golax.InterceptorError)

	home.Build(api.Root, articles_dao)

	// Connect articles API
	articles.Build(api.Root, articles_dao)

	// Connect sessions API
	sessions.Build(api.Root, sessions_dao)

	// Connect users API
	users.Build(api.Root, users_dao)

	// Connect login API
	login.Build(api.Root, users_dao)

	return api
}
