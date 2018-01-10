package config

import (
	"blogo/sessions"
	"fmt"

	"blogo/users"

	"github.com/fulldump/goconfig"
)

type Config struct {
	MongoUri string `usage:"MongoDB URI service"`
	HttpAddr string `usage:"TCP port to listen"`
	Cookies  Cookies
	Users    Users
}

type Cookies struct {
	SecretSalt string `usage:"Secret salt for cookies hashing"`
}

type Users struct {
	SecretSalt string `usage:"Secret salt for user password hashing"`
}

func Read() *Config {

	c := &Config{
		MongoUri: "mongodb://localhost:27017/blogo",
	}

	goconfig.Read(c)

	if c.Cookies.SecretSalt == "" {
		// TODO: use a logger insted of stdout
		fmt.Println("WARNING: default cookies salt used, not ready for production")
	} else {
		sessions.SECRET_SALT = c.Cookies.SecretSalt
	}

	if c.Users.SecretSalt == "" {
		// TODO: use a logger insted of stdout
		fmt.Println("WARNING: default password salt used, not ready for production")
	} else {
		users.SECRET_SALT = c.Users.SecretSalt
	}

	return c
}
