package goaudit

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
)

func Test_Audit(t *testing.T) {

	// Test configuration
	method := "POST"
	url := "/value-1/value-2/test-node?query_a=aaa&query_b=bbb"
	authID := "my-auth-id"
	requestBody := "hi"
	responseBody := "hello"

	auditTest := NewTestAudit()

	service := &Service{
		Name:    "invented-service",
		Version: "7.3.0",
		Commit:  "70bbda5",
	}

	a := golax.NewApi()
	a.Root.
		Interceptor(auditTest.InterceptorAudit2Memory()).
		Interceptor(InterceptorAudit2Log()).
		Interceptor(InterceptorAudit(service)).
		Node("{param1}").
		Node("{param2}").
		Node("test-node").
		Method("POST", func(c *golax.Context) {

			audit := GetAudit(c)
			audit.AuthId = authID
			audit.Custom = map[string]interface{}{
				"a": 20,
				"b": 55,
			}
			audit.AddReadAccess(authID)
			audit.AddReadAccess("other-involved-auth-id")

			c.Error(222, "my-error-description").ErrorCode = 27

			fmt.Fprintf(c.Response, responseBody)
		})

	s := apitest.New(a)

	s.Request(method, url).
		WithBodyString(requestBody).
		Do()

	audit := auditTest.Memory[0]

	if authID != audit.AuthId {
		t.Error("authID", authID, audit.AuthId)
	}

	if nil == audit.Service {
		t.Error("service")
	}

	if "invented-service" != audit.Service.Name {
		t.Error("service.name")
	}

	if "7.3.0" != audit.Service.Version {
		t.Error("service.version")
	}

	if "70bbda5" != audit.Service.Commit {
		t.Error("service.commit")
	}

	if method != audit.Request.Method {
		t.Error("request.method")
	}

	if url != audit.Request.URI {
		t.Error("request.uri")
	}

	if "/{param1}/{param2}/test-node" != audit.Request.Handler {
		t.Error("request.handler")
	}

	if int64(len(requestBody)) != audit.Request.Length {
		t.Error("request.length")
	}

	query := map[string][]string{
		"query_a": []string{"aaa"},
		"query_b": []string{"bbb"},
	}

	if !reflect.DeepEqual(query, audit.Request.Query) {
		t.Error("request.query")
	}

	parameters := map[string]string{
		"param1": "value-1",
		"param2": "value-2",
	}

	if !reflect.DeepEqual(parameters, audit.Request.Parameters) {
		t.Error("request.parameters")
	}

	if int64(len(responseBody)) != audit.Response.Length {
		t.Error("response.length")
	}

	if 222 != audit.Response.StatusCode {
		t.Error("response.status_code")
	}

	if 27 != audit.Response.Error.Code {
		t.Error("response.error.code")
	}

	if "my-error-description" != audit.Response.Error.Description {
		t.Error("response.error.description")
	}

	access := []string{authID, "other-involved-auth-id"}
	if !reflect.DeepEqual(access, audit.ReadAccess) {
		t.Error("read_access", access, audit.ReadAccess)
	}

}
