package app

import (
	"context"
	"github.com/core-go/health"
	mgo "github.com/core-go/health/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handler"
	"go-service/internal/service"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   *handler.UserHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(root.Mongo.Database)

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	mongoChecker := mgo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
