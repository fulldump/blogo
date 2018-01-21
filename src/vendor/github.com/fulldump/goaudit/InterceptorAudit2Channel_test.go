package goaudit

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
)

func Test_Audit2channel(t *testing.T) {

	audits := make(chan *Audit, 10)

	auditTest := NewTestAudit()

	a := golax.NewApi()

	a.Root.
		Interceptor(auditTest.InterceptorAudit2Memory()).
		Interceptor(InterceptorAudit2Channel(audits)).
		Interceptor(InterceptorAudit(&Service{Name: "invented-service"})).
		Node("api").
		Method("GET", func(c *golax.Context) {

			GetAudit(c).Custom = map[string]interface{}{
				"name": c.Request.URL.Query().Get("name"),
			}

		})

	s := apitest.New(a)

	// Do sample SYNCed requests
	for i := 0; i < 10; i++ {
		s.Request("GET", "/api?name="+strconv.Itoa(i)).Do()
	}

	// Extract from channel and compare to memory
	for _, memoryAudit := range auditTest.Memory {
		channelAudit := <-audits
		if !reflect.DeepEqual(channelAudit, memoryAudit) {
			t.Error("Channel Audit does not match with memory Audit")
		}
	}

}
