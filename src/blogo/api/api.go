package api

import (
	"blogo/articles"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func Build(articles_dao *kip.Dao) *golax.Api {

	api := golax.NewApi()

	// Connect articles API
	articles.Build(api.Root, articles_dao)

	return api
}
