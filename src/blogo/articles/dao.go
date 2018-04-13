package articles

import (
	"github.com/fulldump/kip"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

	kip.Define(&kip.Collection{
		Name: "articles",
		OnCreate: func() interface{} {
			return &Article{
				Id: bson.NewObjectId().Hex(),
			}
		},
		Indexes: []mgo.Index{
			{
				Name:       "Unique-articles",
				Key:        []string{"user._id", "title_url"},
				Unique:     true,
				DropDups:   true,
				Sparse:     true,
				Background: true,
			},
		},
	})

}
