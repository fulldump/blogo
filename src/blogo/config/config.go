package config

import (
	"blogo/sessions"
	"fmt"

	"googleapi"

	"github.com/fulldump/goconfig"

	"blogo/constants"
	"blogo/users"
	"os"
)

type Config struct {
	MongoUri        string `usage:"MongoDB URI service"`
	HttpAddr        string `usage:"TCP port to listen"`
	Cookies         Cookies
	Users           Users
	Google          googleapi.GoogleApi
	GoogleAnalytics string `usage:"Google Analytics ID"`

	Version bool `usage:"Show version"`
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
		HttpAddr: ":8000",
	}

	goconfig.Read(c)

	if c.Version {
		fmt.Println("Version:", constants.VERSION, "\tCompiler:", constants.COMPILER)
		os.Exit(0)
	}

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
