package apidoc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
)

func Test_apidoc_happyPath(t *testing.T) {

	a := golax.NewApi()

	// Implement my api
	a.Root.Node("my-api").Method("GET", func(c *golax.Context) {
		// Do nothing
	}, golax.Doc{
		Name:        `Feature`,
		Description: `Feature description`,
	})

	// Add custom endpoint to mount apidoc
	myendpoint := a.Root.Node("a").Node("b").Node("c")

	// Build apidoc on custom endpoint
	d := Build(a, myendpoint)
	d.Title = `My custom title`
	d.Subtitle = `My custom subtitle`

	// Test
	at := apitest.New(a)

	r := at.Request("GET", "/a/b/c/doc/json").Do()

	expected := map[string]interface{}{
		"title":    "My custom title",
		"subtitle": "My custom subtitle",
		"endpoints": map[string]interface{}{
			"": map[string]interface{}{

				"description":  "",
				"interceptors": []interface{}{},
				"methods":      map[string]interface{}{},
			},
			"/a": map[string]interface{}{
				"description":  "",
				"interceptors": []interface{}{},
				"methods":      map[string]interface{}{},
			},
			"/a/b": map[string]interface{}{
				"description":  "",
				"interceptors": []interface{}{},
				"methods":      map[string]interface{}{},
			},
			"/a/b/c": map[string]interface{}{
				"description":  "",
				"interceptors": []interface{}{},
				"methods":      map[string]interface{}{},
			},
			"/my-api": map[string]interface{}{
				"description":  "",
				"interceptors": []interface{}{},
				"methods": map[string]interface{}{
					"GET": map[string]interface{}{
						"name":        "Feature",
						"description": "Feature description",
					},
				},
			},
		},
		"interceptors": map[string]interface{}{},
	}

	body := r.BodyJson()

	if !reflect.DeepEqual(body, expected) {
		t.Error("Body json does not match")
		fmt.Println(body)
	}

}
