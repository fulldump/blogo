package main

import (
	"blogo/config"

	"blogo/api"

	"github.com/fulldump/kip"
)

func main() {

	// Get config
	c := config.Read()

	// Connect to Mongo
	db, db_err := kip.NewDatabase(c.MongoUri)
	if nil != db_err {
		panic(db_err)
	}

	articles_dao := kip.NewDao("articles", db)
	sessions_dao := kip.NewDao("sessions", db)
	users_dao := kip.NewDao("users", db)

	// Buid API
	a := api.Build(articles_dao, sessions_dao, users_dao)

	// Serve
	a.Serve()
}
