package goaudit

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Test_Chan2Mongo(t *testing.T) {

	dbName := "goaudit-" + uuid.Must(uuid.NewV4()).String()

	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	session.SetSyncTimeout(100 * time.Millisecond) // Insert fast fail

	db := session.DB(dbName)

	defer func(d *mgo.Database) {
		d.DropDatabase()
	}(db)

	audits := make(chan *Audit, 10) // #1 make channel

	Chan2Mongo(audits, db) // #2 dump channel to mongo

	auditTest := NewTestAudit()

	a := golax.NewApi()
	a.Root.
		Interceptor(auditTest.InterceptorAudit2Memory()).
		Interceptor(InterceptorAudit2Channel(audits)). // #3 put Audits in channel
		Interceptor(InterceptorAudit(nil)).
		Node("api").
		Method("GET", func(c *golax.Context) {})

	s := apitest.New(a)

	// Do sample SYNCed requests
	for i := 0; i < 10; i++ {
		s.Request("GET", "/api?name="+strconv.Itoa(i)).Do()
	}

	// wait all audits to be stored in mongo
	for len(audits) != 0 {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("chan.length:", len(audits))
	}

	// Extract from mongo and compare to memory
	for i, memoryAudit := range auditTest.Memory {

		collection := db.C(getCollectionName(memoryAudit))

		mongoAudit := &Audit{}
		err := collection.Find(bson.M{"request.query.name": strconv.Itoa(i)}).One(mongoAudit)
		if nil != err {
			panic(err)
		}

		// Tweak differences between mongo representation and memory:
		mongoAudit.Id = memoryAudit.Id
		mongoAudit.Custom = map[string]interface{}{}
		mongoAudit.EntryDate = memoryAudit.EntryDate

		if !reflect.DeepEqual(mongoAudit, memoryAudit) {
			t.Error("Channel Audit does not match with memory Audit")
		}
	}

}
