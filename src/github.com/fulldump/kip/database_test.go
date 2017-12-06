package kip

import (
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DATABASE_HOST       = "localhost"
	DATABASE_NAME       = random_name("database")
	DATABASE_COLLECTION = random_name("collection")
)

func openTestDatabase() (*Database, error) {
	return NewDatabase("mongodb://" + DATABASE_HOST + "/" + DATABASE_NAME)
}

func closeTestDatabase() {
	session, _ := mgo.Dial(DATABASE_HOST)
	session.SetMode(mgo.Monotonic, true)
	session.DB(DATABASE_NAME).DropDatabase()
	session.Close()

	CloseAll()
}

func TestDatabase_Configure_TwoDBs_OneHost(t *testing.T) {
	host := "localhost"
	name1 := random_name("db1")
	name2 := random_name("db2")

	db1, _ := NewDatabase(host + "/" + name1)
	db2, _ := NewDatabase(host + "/" + name2)

	if db1.session != db2.session {
		t.Error("Two dbs with the same host should have the same session")
	}
}

func TestDatabase_Configure_TwoDBs_TwoHosts(t *testing.T) {
	host1 := "localhost"
	host2 := "127.0.0.1"
	name1 := random_name("db1")
	name2 := random_name("db2")

	db1, _ := NewDatabase(host1 + "/" + name1)
	db2, _ := NewDatabase(host2 + "/" + name2)

	if db1.session == db2.session {
		t.Error("Two dbs with the different hosts should have different sessions")
	}
}

func TestDatabase_ConfigureDatabase_HappyPath(t *testing.T) {
	_, err := openTestDatabase()
	defer closeTestDatabase()

	if nil == err {
		t.Log("Connection to mongo stablished ok")
	} else {
		t.Skip("Fail to connect to mongo. Ensure there is a mongo runnin in 'localhost' ")
	}
}

func TestDatabase_GetDatabase_HappyPath(t *testing.T) {
	database, _ := openTestDatabase()
	defer closeTestDatabase()

	err := database.session.Ping()

	if nil == err {
		t.Log("Connection seems to be working, ping received")
	} else {
		t.Error("Connection is not working")
	}
}

func TestDatabase_C_HappyPath(t *testing.T) {
	database, _ := openTestDatabase()
	defer closeTestDatabase()

	collection := database.C(DATABASE_COLLECTION)

	// Data set
	_id := bson.NewObjectId()
	name := "test name 0123456789"

	// Insert the object
	err := collection.Insert(bson.M{
		"_id":  _id,
		"Name": name,
	})
	if nil != err {
		t.Error("Fail inserting document")
	}

	// Retrieve the object
	result := struct {
		Name string `bson:"Name"`
	}{}

	err = collection.Find(bson.M{"_id": _id}).One(&result)
	if nil != err {
		t.Error("Fail retrieving the document")
	}

	// Conclusion
	if result.Name == name {
		t.Log("The collection works fine")
	} else {
		t.Log("The collection does not work")
	}
}
