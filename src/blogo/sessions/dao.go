package sessions

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/fulldump/kip"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var SECRET_SALT = "227a61ee-f309-11e7-a02d-a340fcd61bb3"

func init() {

	kip.Define(&kip.Collection{
		Name: "sessions",
		OnCreate: func() interface{} {

			now := time.Now()

			return &Session{
				Id:              bson.NewObjectId().Hex(),
				CreateTimestamp: now.Unix(),
				ExpireTimestamp: now.Add(24 * time.Hour).Unix(),
				Data:            map[string]interface{}{},
			}

		},
		Indexes: []mgo.Index{
			mgo.Index{
				Key:        []string{"-create_timestamp"},
				Unique:     false,
				Sparse:     false,
				Background: false,
				DropDups:   false,
			},
			mgo.Index{
				Key:        []string{"cookie"},
				Unique:     true,
				Sparse:     false,
				Background: false,
				DropDups:   true,
			},
		},
	})

}

func NewSession(dao *kip.Dao) (cookie string, item *kip.Item) {

	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	cookie = u.String()
	item = dao.Create()

	session := item.Value.(*Session)
	session.Cookie = hash(cookie)

	return
}

func hash(i string) (o string) {

	adder := md5.New()
	adder.Write([]byte(SECRET_SALT))
	adder.Write([]byte(i))

	return hex.EncodeToString(adder.Sum(nil))
}
