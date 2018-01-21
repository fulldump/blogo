package goaudit

import (
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
)

// TestAudit is a util for testing Audits. It can store Audits to memory and read
// them at testing time.
type TestAudit struct {

	// Memory is the secuence of Audits in a test
	Memory []*Audit
}

// Reset clean all Audits from memory. SetUp test is a good place to use it.
func (t *TestAudit) Reset() {
	t.Memory = []*Audit{}
}

// NewTestAudit Create a new TestAudit
//
// Typical usage:
//
// 	func Test_ExampleBilling(t *testing.T) {
//
// 	    audittest := testutils.NewTestAudit() // IMPORTANT
//
// 	    a := golax.NewApi()
// 	    a.Root.Interceptor(audittest.InterceptorAudit2Memory()) // IMPORTANT
//
// 	    BuildYourApi(a)
//
// 	    s := apitest.New(a)
// 	    s.Request("GET", "/my-url-to-test").Do()
//
// 	    // IMPORTANT: Do things with audittest.Memory[i], for example:
// 	    if 200 != audittest.Memory[0].Response.StatusCode {
// 	        t.Error("blah blah blah...")
// 	    }
// 	}
func NewTestAudit() *TestAudit {
	t := &TestAudit{}
	t.Reset()

	return t
}

// InterceptorAuit2Memory return the interceptor you should put in your api to
// capture Audits
func (t *TestAudit) InterceptorAudit2Memory() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name: "Audit2Memory",
			Description: `
			Save audits into memory.
		`,
		},
		After: func(c *golax.Context) {
			audit := GetAudit(c)
			t.Memory = append(t.Memory, audit)
		},
	}
}

func Test_Audit2memory(t *testing.T) {

	audittest := NewTestAudit()

	a := golax.NewApi()

	a.Root.
		Interceptor(audittest.InterceptorAudit2Memory()).
		Interceptor(InterceptorAudit2Log()).
		Interceptor(InterceptorAudit(nil)).
		Node("api").
		Method("GET", func(c *golax.Context) {

			name := c.Request.URL.Query().Get("name")

			audit := GetAudit(c)
			audit.Custom = map[string]interface{}{
				"name": name,
			}

		})

	s := apitest.New(a)

	method := "GET"
	url := "/api?name="

	s.Request(method, url+"one").Do()

	s.Request(method, url+"two").Do()

	if audittest.Memory[0].Custom.(map[string]interface{})["name"] != "one" {
		t.Error("Audit 'one' not found in memory (first position)")
	}

	if audittest.Memory[1].Custom.(map[string]interface{})["name"] != "two" {
		t.Error("Audit 'two' not found in memory (second position)")
	}

	// Check Reset()
	audittest.Reset()
	if !reflect.DeepEqual(audittest.Memory, []*Audit{}) {
		t.Error("`TestAudit.Reset()` should empty the `Memory` array.")
		return
	}

}
