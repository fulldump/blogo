package goaudit

import (
	"encoding/json"
	"log"

	"github.com/fulldump/golax"
)

// InterceptorAudit2Log log audits to `stdout`. This interceptor should wrap
// `InterceptorAudit`.
//
// Typical usage:
//
// 	a := golax.NewApi()
//
// 	a.Root.
// 	    Interceptor(InterceptorAudit2Log()).
// 	    Interceptor(InterceptorAudit(&Service{Name: "invented-service"})).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
//
// Here is the sample output:
//
// 	2016/10/24 19:49:19 AUDIT {"id":"580a465bce507629a613107c","version":"1.0.0","auth_id":"my-auth-id","origin":"127.0.0.1","session_id":"","service":"invented-service","entry_date":"2016-10-21T18:46:19.820423299+02:00","entry_timestamp":1.4770683798204234e+09,"elapsed_seconds":1.621246337890625e-05,"request":{"method":"POST","uri":"/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb","handler":"/{param1}/{param2}/test-node","args":{"query_a":["aaa"],"query_b":["bbb"]},"length":2},"response":{"status_code":222,"length":5,"error":{"code":27,"description":"my-error-description"}},"read_access":["other-involved-auth-id","my-auth-id"],"custom":{"a":20,"b":55}}
//
func InterceptorAudit2Log() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Audit2Log",
			Description: `
			Store Audits into log.
		`,
		},
		After: func(c *golax.Context) {

			audit := GetAudit(c)

			serialized, err := json.Marshal(audit)

			if nil != err {
				// TODO: Log this somewere
			}

			log.Printf("%s\t%s", "AUDIT", serialized)
		},
	}
}
