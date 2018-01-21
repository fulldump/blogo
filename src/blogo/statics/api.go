package statics

import (
	"github.com/fulldump/golax"
)

func Build(node *golax.Node, statics string) {

	m := readFileInternal
	if "" != statics {
		m = readFileExternal(statics)
	}

	node.Node("{{*}}").Method("GET", func(c *golax.Context) {
		m(c, c.Parameter)
	})

}
