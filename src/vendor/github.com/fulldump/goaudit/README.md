# goaudit

[![Build Status](https://travis-ci.org/fulldump/goaudit.svg?branch=master)](https://travis-ci.org/fulldump/goaudit)
[![Go report card](http://goreportcard.com/badge/fulldump/goaudit)](https://goreportcard.com/report/fulldump/goaudit)
[![GoDoc](https://godoc.org/github.com/fulldump/goaudit?status.svg)](https://godoc.org/github.com/fulldump/goaudit)

<sup>Tested for Go 1.5, 1.6, 1.7, tip</sup>

Audits SDK

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [InterceptorAudit](#interceptoraudit)
	- [Custom info](#custom-info)
	- [Access](#access)
	- [Sample audit](#sample-audit)
- [InterceptorAudit2Log](#interceptoraudit2log)
- [InterceptorAudit2Channel](#interceptoraudit2channel)
- [GetAudit](#getaudit)
- [NewTestAudit](#newtestaudit)
- [Dependencies](#dependencies)
- [Testing](#testing)

<!-- /MarkdownTOC -->


Goaudit is a set of interceptors to generate and store audits:

* [InterceptorAudit](#interceptoraudit) - Generate and fill an Audit
* [InterceptorAudit2Log](#interceptoraudit2log) - Log Audits
* [InterceptorAudit2Channel](#interceptoraudit2channel) - Collect all Audits in a channel

And some helpers:

* [GetAudit](#getaudit) - Retrieve Audit from context
* [NewTestAudit](#newtestaudit) - Store Audits in memory to allow unit testing


## InterceptorAudit

Attach a new audit to the context and populate it automatically.

Typical usage:

```go
a := golax.NewApi()

a.Root.
	Interceptor(goaudit.InterceptorAudit2Log()).
	Interceptor(goaudit.InterceptorAudit(&goaudit.Service{Name:"invented-service"})).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})

```

### Custom info

Audit allow to store custom service information:

```go
audit := GetAudit(c)
audit.Custom = map[string]interface{}{
	"a": 20,
	"b": 55,
}
```

### Access

Typically only the user using the service can read the audit but sometimes, a
third party is involved (for example, somebody interacting with me).

In that case, the service should add the third auth_id to the access list:

```go
audit := GetAudit(c)
audit.AddReadAccess("other-involved-auth-id")
```

### Sample audit


A sample audit in JSON:
```json
{
	"id": "580a465bce507629a613107c",
	"version": "2.0.0",
	"auth_id": "my-auth-id",
	"origin": "127.0.0.1",
	"session_id": "",
	"service": {
		"name": "invented-service"
	},
	"entry_date": "2016-10-21T18:46:19.820423299+02:00",
	"entry_timestamp": 1.4770683798204234e+09,
	"elapsed_seconds": 1.621246337890625e-05,
	"request": {
		"method": "POST",
		"uri": "/value-1/value-2/test-node?query_a=aaa\u0026query_b=bbb",
		"handler": "/{param1}/{param2}/test-node",
		"args": {
			"query_a": ["aaa"],
			"query_b": ["bbb"]
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
```


```
TODO: Explain all fields in an Audit
```

## InterceptorAudit2Log

Log Audit to `stdout`. This interceptor should wrap `InterceptorAuth`.

Typical usage:

```go
a := golax.NewApi()

a.Root.
	Interceptor(goaudit.InterceptorAudit2Log()).
	Interceptor(goaudit.InterceptorAudit(&goaudit.Service{Name:"invented-service"})).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})

```



## InterceptorAudit2Channel

Push Audits to a channel. This interceptor should wrap `InterceptorAudit`.

Note that this interceptor do not extract the Audits from the channel, that should
be done by other task. There is a util included (`Chan2Mongo`) to do that.

Typical usage:
```go
channel_audits := make(chan *goaudit.Audit, 100)  // Buffered channel, 100 items

// assume `mongo_db` already exists

Chan2Mongo(channel_audits, mongo_db) // do the job: channel -> mongo

a := golax.NewApi()

a.Root.
	Interceptor(goaudit.InterceptorAudit2Channel(channel_audits)). // Pass created channel
	Interceptor(goaudit.InterceptorAudit(&goaudit.Service{Name:"invented-service"})).
	Method("GET", func(c *golax.Context) {
		// Implement your API here
	})
```

## GetAudit

Get Audit object from context.

Typical usage:

```go
func MyHandler(c *golax.Context) {
	// ...
	audit := goaudit.GetAudit(c)
	// ...
}
```

## NewTestAudit

Store Audits in memory to be readed in your tests.

Typical usage:

```go
func Test_ExampleBilling(t *testing.T) {

	audittest := testutils.NewTestAudit() // IMPORTANT

	a := golax.NewApi()
	a.Root.Interceptor(audittest.InterceptorAudit2Memory()) // IMPORTANT

	BuildYourApi(a)

	s := apitest.New(a)
	s.Request("GET", "/my-url-to-test").Do()

	// IMPORTANT: Do things with audittest.Memory[i], for example:
	if 200 != audittest.Memory[0].Response.StatusCode {
		t.Error("blah blah blah...")
	}
}
```

## Dependencies

Dependencies for testing are:

* github.com/fulldump/apitest
* github.com/fulldump/golax
* github.com/satori/go.uuid
* gopkg.in/mgo.v2

NOTE: Pinned versions are NOT included in `vendor/*`.

Transitive dependencies for runtime are:

* github.com/fulldump/golax

If `goaudit.Chan2Mongo` is used, as you can expect, it will be needed also:

* gopkg.in/mgo.v2


## Testing

As simple as:

```sh
git clone "<this-repo>"
make
```
