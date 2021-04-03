package app

import (
	"context"

	"github.com/common-go/health"
	"github.com/common-go/mongo"

	"go-service/internal/handlers"
	"go-service/internal/services"
)

type ApplicationContext struct {
	HealthHandler *health.HealthHandler
	UserHandler   *handlers.UserHandler
}

func NewApp(ctx context.Context, mongoConfig mongo.MongoConfig) (*ApplicationContext, error) {
	db, err := mongo.SetupMongo(ctx, mongoConfig)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	mongoChecker := mongo.NewHealthChecker(db)
	checkers := []health.HealthChecker{mongoChecker}
	healthHandler := health.NewHealthHandler(checkers)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
