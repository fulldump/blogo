package goaudit

import (
	"net/http"
	"strings"
	"time"

	"github.com/fulldump/golax"
	"gopkg.in/mgo.v2/bson"
)

// InterceptorAudit attach a new Audit to the context and populate it automatically.
//
// Typical usage:
//
// 	a := golax.NewApi()
//
// 	s := &model.Service{
// 	    Name: "My service",
// 	    Version: "7.3.0",
// 	    Commit: "0f0710f",
// 	}
//
// 	a.Root.
// 	    Interceptor(InterceptorAudit2Log()).
// 	    Interceptor(InterceptorAudit(s)).
// 	    Method("GET", func(c *golax.Context) {
// 	        // Implement your API here
// 	    })
func InterceptorAudit(service *Service) *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "InterceptorAudit",
			Description: `
			Collect standard request information and custom service information
			and makeis available for processing it.

			Sample Audit:
			´´´json

			{
				"id": "580a465bce507629a613107c",
				"version": "1.0.0",
				"auth_id": "my-auth-id",
				"session_id": "",
				"origin": "127.0.0.1",
				"service": {
					"name": "invented-service",
					"version": "7.3.0",
					"commit": "a587df8"
				},
				"entry_date": "2016-10-21T18:46:19.820423299+02:00",
				"entry_timestamp": 1.4770683798204234e+09,
				"elapsed_seconds": 1.621246337890625e-05,
				"request": {
					"method": "POST",
					"uri": "/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb",
					"handler": "/{param1}/{param2}/test-node",
					"query": {
						"query_a": ["aaa"],
						"query_b": ["bbb"]
					},
					"parameters": {
						"param1": "value1",
						"param2": "value2"
					},
					"length": 2
				},
				"response": {
					"status_code": 222,
					"length": 5,
					"error": {
						"code": 27,
						"description": "my-error-description"
					}
				},
				"read_access": ["other-involved-auth-id", "my-auth-id"],
				"custom": {
					"a": 20,
					"b": 55
				}
			}
			´´´
		`,
		},
		Before: func(c *golax.Context) {

			audit := &Audit{
				Id:             bson.NewObjectId().Hex(),
				Version:        VERSION,
				Service:        service,
				Origin:         formatRemoteAddr(c.Request),
				EntryDate:      time.Now(),
				EntryTimestamp: float64(time.Now().UnixNano()) / 1000000000,
				ElapsedSeconds: 0,
				Request: Request{
					Header:     c.Request.Header,
					Length:     c.Request.ContentLength,
					Method:     c.Request.Method,
					URI:        c.Request.RequestURI,
					Query:      c.Request.URL.Query(),
					Parameters: c.Parameters,
				},
				ReadAccess: []string{},
				Custom:     map[string]interface{}{},
				Log: Log{
					Entries: []*LogEntry{},
				},
			}

			c.Response.Header().Set("X-Audit-Id", audit.Id)

			c.Set(CONTEXT_KEY, audit)

		},
		After: func(c *golax.Context) {

			audit := GetAudit(c)

			exitTimestamp := float64(time.Now().UnixNano()) / 1000000000
			audit.ElapsedSeconds = exitTimestamp - audit.EntryTimestamp

			audit.Request.Handler = c.PathHandlers
			audit.Response.Header = c.Response.Header()
			audit.Response.StatusCode = c.Response.StatusCode
			audit.Response.Length = int64(c.Response.Length)

			err := c.LastError
			if nil != err {
				audit.SetError(err.ErrorCode, err.Description)
			}

		},
	}
}

func formatRemoteAddr(r *http.Request) string {
	xorigin := strings.TrimSpace(strings.Split(
		r.Header.Get("X-Forwarded-For"), ",")[0])
	if xorigin != "" {
		return xorigin
	}

	return r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
}
