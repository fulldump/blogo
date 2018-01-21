package kip

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Email   string        `bson:"email"`
	Age     int           `bson:"age"`
	Single  bool          `bson:"single"`
	Friends []string      `bson:"friends"`
	Colors  []string      `bson:"colors"`
	Lock    bool          `bson:"lock"`
}
