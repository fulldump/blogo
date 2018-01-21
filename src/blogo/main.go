package main

import (
	"blogo/config"

	"blogo/api"

	"fmt"
	"net/http"

	"github.com/fulldump/goaudit"
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
	audits_dao := kip.NewDao("audits", db)

	// audits channel
	channel_audits := make(chan *goaudit.Audit, 1000000)  // Buffered channel, 100 items
	goaudit.Chan2Mongo(channel_audits, db.C("").Database) // do the job: channel -> mongo

	// Buid API
	a := api.Build(articles_dao, sessions_dao, users_dao, audits_dao, &c.Google, c.Statics, c.GoogleAnalytics, channel_audits)

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
