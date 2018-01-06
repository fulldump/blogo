package config

import (
	"blogo/sessions"
	"fmt"

	"github.com/fulldump/goconfig"
)

type Config struct {
	MongoUri   string `usage:"MongoDB URI service"`
	HttpAddr   string `usage:"TCP port to listen"`
	SecretSalt string `usage:"Secret salt for cookies hashing"`
}

func Read() *Config {

	c := &Config{
		MongoUri: "mongodb://localhost:27017/blogo",
	}

	goconfig.Read(c)

	if c.SecretSalt == "" {
		// TODO: use a logger insted of stdout
		fmt.Println("WARNING: default salt used, not ready for production")
	} else {
		sessions.SECRET_SALT = c.SecretSalt
	}

	return c
}
