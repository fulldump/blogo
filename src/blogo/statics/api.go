package statics

import (
	"fmt"

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
		m(c, c.Parameter)
	})

}
