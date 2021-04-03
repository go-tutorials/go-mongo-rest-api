package app

import (
	"github.com/common-go/log"
	m "github.com/common-go/middleware"
)

type Root struct {
	Server     ServerConfig `mapstructure:"server"`
	Mongo      MongoConfig  `mapstructure:"mongo"`
	Log        log.Config   `mapstructure:"log"`
	MiddleWare m.LogConfig  `mapstructure:"middleware"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}
