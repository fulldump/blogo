package config

import "github.com/fulldump/goconfig"

type Config struct {
	MongoUri string `usage:"MongoDB URI service"`
	HttpAddr string `usage:"TCP port to listen"`
}

func Read() *Config {

	c := &Config{
		MongoUri: "mongodb://localhost:27017/blogo",
	}

	goconfig.Read(c)

	return c
}