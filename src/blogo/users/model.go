package users

import (
	"crypto/md5"
	"encoding/hex"
)

var SECRET_SALT = "fb30b9bc-f3e9-11e7-ae16-0369d9f380f8"

type User struct {
	Id              string     `bson:"_id" json:"id"`
	Nick            string     `bson:"nick" json:"nick"`
	CreateTimestamp int64      `bson:"create_timestamp" json:"create_timestamp"`
	LoginEmail      LoginEmail `bson:"login_email" json:"login_email"`
	Scopes          Scopes     `bson:"scopes" json:"scopes"`
}

type LoginEmail struct {
	Email        string `bson:"email" json:"email"`
	PasswordHash string `bson:"password_hash" json:"password_hash"`
}

func (e LoginEmail) Check(password string) bool {
	password_hash := hash(password)

	return password_hash == e.PasswordHash
}

type Scopes struct {
	Admin  bool `bson:"admin" json:"admin"`
	Banned bool `bson:"banned" json:"banned"`
}

func hash(i string) (o string) {

	adder := md5.New()
	adder.Write([]byte(SECRET_SALT))
	adder.Write([]byte(i))

	return hex.EncodeToString(adder.Sum(nil))
}
