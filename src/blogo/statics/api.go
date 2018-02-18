package statics

import (
	"fmt"

	"strings"

	"github.com/fulldump/golax"
)

func Build(node *golax.Node, statics string) {

	m := readFileInternal

	if "" != statics {
		m = readFileExternal(statics)
	}

	node.Node("statics").Method("GET", func(c *golax.Context) {
		for filename, _ := range Files {
			fmt.Fprintln(c.Response, "<a href='/"+filename+"'>"+filename+"</a><br>")
		}
		for filename, _ := range Bytes {
			fmt.Fprintln(c.Response, "<a href='/"+filename+"'>"+filename+"</a><br>")
		}
	})

	node.Node("{{*}}").Method("GET", func(c *golax.Context) {

		filename := c.Request.URL.Path

		fmt.Println(filename)

		if strings.HasSuffix(filename, "/") {
			filename += "index.html"
		}
		m(c, filename)
	})

}
