package goaudit

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Chan2Mongo is a util to extract Audits from a channel and store to a MongoDB
// database. This interceptor should wrap `InterceptorAudit`.
//
// Note that this interceptor do not extract the Audits from the channel, that should
// be done by other task. There is a util included (`Chan2Mongo`) to do that.
//
// Typical usage:
// 	channel_audits := make(chan *Audit, 100)  // Buffered channel, 100 items
//
// 	assume `mongo_db` already exists
//
// 	Chan2Mongo(channel_audits, mongo_db) // do the job: channel -> mongo
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorAudit2Channel(channel_audits)). // Pass created channel
// 	    Interceptor(InterceptorAudit("invented-service")).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
//
func Chan2Mongo(s chan *Audit, m *mgo.Database) {

	go func() {

		for {
			audit := <-s
			if nil == audit {
				break
			}

			name := getCollectionName(audit)
			c := m.C(name)

			for e := c.Insert(audit); nil != e; e = c.Insert(audit) {
				if mgo.IsDup(e) {
					break
				}
				// TODO: alarm warning?
				time.Sleep(1 * time.Second)
			}

		}

	}()

}

// GetCollectionName generates the name for the collection to split Audit
// across timestamp
func getCollectionName(audit *Audit) string {
	return "audits"
	// return fmt.Sprintf(
	// 	"audit%04d%02d%02d",
	// 	audit.EntryDate.Year(),
	// 	audit.EntryDate.Month(),
	// 	audit.EntryDate.Day(),
	// )
}
