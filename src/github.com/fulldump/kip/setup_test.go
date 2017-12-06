package kip

import (
	"testing"

	. "gopkg.in/check.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type World struct {
	MongoHosts    string
	MongoDatabase string
	Database      *Database
	Kip           *Kip
	Users         *Dao
}

var _ = Suite(&World{
	MongoHosts: "localhost",
})

func (w *World) SetUpSuite(c *C) {

	db, err := NewDatabase(w.MongoHosts + "/" + random_name("kip"))
	if err != nil {
		panic("Fail creating a TESTING database. Please, check your MongoDB")
	}
	w.Database = db

}

func (w *World) SetUpTest(c *C) {

	w.Kip = NewKip()

	w.Kip.Define(&Collection{
		Name: "Users",
		OnCreate: func() interface{} {
			return &User{
				Id:      bson.NewObjectId(),
				Friends: []string{},
				Colors:  []string{},
			}
		},
	})

	w.Users = w.Kip.NewDao("Users", w.Database)
}

func (w *World) TearDownTest(c *C) {

	// When all tests are finished, drop database
	session, _ := mgo.Dial(w.MongoHosts)
	session.SetMode(mgo.Monotonic, true)
	session.DB(w.Database.name).DropDatabase()
	session.Close()
}

func (w *World) TearDownSuite(c *C) {
	//fmt.Println("SLEEPING...")
	//time.Sleep(20 * time.Second)
}
