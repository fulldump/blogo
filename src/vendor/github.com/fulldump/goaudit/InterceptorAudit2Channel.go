package goaudit

import (
	"github.com/fulldump/golax"
)

// InterceptorAudit2Channel push Audits to a channel. This interceptor should wrap
// `InterceptorAudit`.
//
// Note that this interceptor do not extract the Audits from the channel, that should
// be done by other task. There is a util included (`Chan2Mongo`) to do that.
//
// Typical usage:
//
// 	channel_audits := make(chan *Audit, 100)  // Buffered channel, 100 items
//
// 	// assume `mongo_db` already exists
//
// 	Chan2Mongo(channel_audits, mongo_db) // do the job: channel -> mongo
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorAudit2Channel(channel_audits)). // Pass created channel
// 	    Interceptor(InterceptorAudit(&Service{Name: "invented service"}).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
func InterceptorAudit2Channel(s chan *Audit) *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Audit2Channel",
			Description: `
			Push Auditss to a channel.
		`,
		},
		After: func(c *golax.Context) {
			audit := GetAudit(c)
			if audit.aborted {
				return
			}
			s <- audit
		},
	}
}
