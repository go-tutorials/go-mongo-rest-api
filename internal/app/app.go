package app

import (
	"context"
	"go-service/internal/services/client_repository"
	"go-service/internal/services/sync_repository"
	"go-service/internal/services/sync_service"
	"go-service/internal/services/tube_category_service"
	"go-service/internal/services/tube_service"
	"log"

	"github.com/core-go/health"
	mgo "github.com/core-go/health/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handlers"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	SyncHandler   *handlers.SyncHandler
	ClientHandler *handlers.ClientHandler
	TubeHandler   *handlers.TubeHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	log.Println(root.Mongo.Uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(root.Mongo.Database)

	mongoChecker := mgo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	tubeCategory := tube_category_service.NewCategoryTubeService(root.Key)

	tubeService := tube_service.NewTubeService(root.Key)
	tubeHandler := handlers.NewTubeHandler(tubeService)

	channelCollectionName := "channel"
	channelSyncCollectionName := "channelSync"
	playlistCollectionName := "playlist"
	playlistVideoCollectionName := "playlistVideo"
	videoCollectionName := "video"
	categoryCollectionName :="category"
	repo := sync_repository.NewMongoVideoRepository(db, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName)
	syncService := sync_service.NewDefaultSyncService(tubeService, repo)
	syncHandler := handlers.NewSyncHandler(syncService)

	clientService := client_repository.NewMongoVideoService(db, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName, *tubeCategory)
	clientHandler := handlers.NewClientHandler(clientService)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		ClientHandler: clientHandler,
		SyncHandler:   syncHandler,
		TubeHandler:   tubeHandler,
	}, nil
}
