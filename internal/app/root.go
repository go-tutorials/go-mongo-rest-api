package app

import (
	"github.com/common-go/log"
	mid "github.com/common-go/log/middleware"
	"github.com/common-go/mongo"
	sv "github.com/common-go/service"
)

type Root struct {
	Server     sv.ServerConfig   `mapstructure:"server"`
	Mongo      mongo.MongoConfig `mapstructure:"mongo"`
	Log        log.Config        `mapstructure:"log"`
	MiddleWare mid.LogConfig     `mapstructure:"middleware"`
}
