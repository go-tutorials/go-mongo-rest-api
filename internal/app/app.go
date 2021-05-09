package app

import (
	"context"
	"github.com/core-go/health"
	mgo "github.com/core-go/health/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handlers"
	"go-service/internal/services"
)

type ApplicationContext struct {
	HealthHandler *health.HealthHandler
	UserHandler   *handlers.UserHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(root.Mongo.Database)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	mongoChecker := mgo.NewHealthChecker(db)
	healthHandler := health.NewHealthHandler(mongoChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
