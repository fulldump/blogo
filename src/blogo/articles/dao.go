package articles

import (
	"github.com/fulldump/kip"
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
	})

}
