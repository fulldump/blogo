package api

import (
	"blogo/articles"

	"blogo/home"
	"blogo/sessions"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(articles_dao, sessions_dao *kip.Dao) *golax.Api {

	api := golax.NewApi()

	api.Root.Interceptor(sessions.NewSessionsInterceptor(sessions_dao))

	home.Build(api.Root, articles_dao)

	// Connect articles API
	articles.Build(api.Root, articles_dao)

	// Connect sessions API
	sessions.Build(api.Root, sessions_dao)

	return api
}
