package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Config struct {
	Server     ServerConfig  `mapstructure:"server"`
	Mongo      MongoConfig   `mapstructure:"mongo"`
	Log        log.Config    `mapstructure:"log"`
	MiddleWare mid.LogConfig `mapstructure:"middleware"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port *int64 `mapstructure:"port"`
}

type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}
