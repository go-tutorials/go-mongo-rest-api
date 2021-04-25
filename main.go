package main

import (
	"context"
	"fmt"
	"github.com/common-go/config"
	"github.com/common-go/log"
	mid "github.com/common-go/log/middleware"
	sv "github.com/common-go/service"
	"github.com/gorilla/mux"
	"net/http"

	"go-service/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewStructuredLogger()
	r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	r.Use(mid.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println(sv.ServerInfo(conf.Server))
	http.ListenAndServe(sv.Addr(conf.Server.Port), r)
}
