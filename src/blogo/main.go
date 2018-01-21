package main

import (
	"blogo/config"

	"blogo/api"

	"fmt"
	"net/http"

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
	a := api.Build(articles_dao, sessions_dao, users_dao, &c.Google, c.GoogleAnalytics)

	// Serve
	s := &http.Server{
		Addr:    c.HttpAddr,
		Handler: a,
	}

	fmt.Println("Listening", s.Addr)
	if err := s.ListenAndServe(); nil != err {
		fmt.Println("Server:", err)
	}

}
