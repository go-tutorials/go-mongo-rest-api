package app

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handlers"
	"go-service/internal/services"
)

type ApplicationContext struct {
	UserHandler     *handlers.UserHandler
}

func NewApp(context context.Context, mongoConfig MongoConfig) (*ApplicationContext, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(mongoConfig.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(mongoConfig.Database)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	return &ApplicationContext{
		UserHandler: userHandler,
	}, nil
}
