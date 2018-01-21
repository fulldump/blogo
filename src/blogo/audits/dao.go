package audits

import (
	"github.com/fulldump/goaudit"
	"github.com/fulldump/kip"
	mgo "gopkg.in/mgo.v2"
)

const NAME = "audits"

func init() {

	kip.Define(&kip.Collection{
		Name: NAME,
		OnCreate: func() interface{} {
			return &goaudit.Audit{}
		},
		Indexes: []mgo.Index{
			mgo.Index{
				Key:        []string{"-entry_timestamp"},
				Unique:     false,
				DropDups:   true,
				Background: false, // See notes.
				Sparse:     false,
			},
		},
	})

}
