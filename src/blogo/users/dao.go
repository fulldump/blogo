package users

import (
	"time"

	"github.com/fulldump/kip"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

	kip.Define(&kip.Collection{
		Name: "users",
		OnCreate: func() interface{} {
			return &User{
				Id:              bson.NewObjectId().Hex(),
				CreateTimestamp: time.Now().Unix(),
				Scopes: Scopes{
					Admin:  false,
					Banned: false,
				},
			}
		},
		Indexes: []mgo.Index{
			mgo.Index{
				Key:        []string{"nick"},
				Unique:     true,
				Sparse:     false,
				Background: false,
				DropDups:   true,
			},
			mgo.Index{
				Key:        []string{"login_email.email"},
				Unique:     true,
				Sparse:     false,
				Background: false,
				DropDups:   true,
			},
		},
	})

}
