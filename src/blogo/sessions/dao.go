package sessions

import (
	"time"

	"github.com/fulldump/kip"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
)

func init() {

	kip.Define(&kip.Collection{
		Name: "sessions",
		OnCreate: func() interface{} {
			return &Session{
				Id:                  bson.NewObjectId().Hex(),
				Cookie:              uuid.NewV4().String(),
				UserId:              "",
				CreateTimestamp:     time.Now().Unix(),
				ExpirationTimestamp: time.Now().Unix() + 24*3600,
			}
		},
		Indexes: []mgo.Index{
			mgo.Index{
				Name:       "cookie_unique",
				Key:        []string{"cookie"},
				Background: false,
				Unique:     true,
				Sparse:     false,
			},
		},
	})

}
