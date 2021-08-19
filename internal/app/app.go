package app

import (
	"context"
	"database/sql"
	"go-service/internal/services/client_repository"
	"go-service/internal/services/sync_repository"
	"go-service/internal/services/sync_service"
	"go-service/internal/services/tube_category_service"
	"go-service/internal/services/tube_service"

	"github.com/core-go/health"
	//mgo "github.com/core-go/health/mongo"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

	s "github.com/core-go/health/sql"
	_ "github.com/lib/pq"

	"go-service/internal/handlers"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	SyncHandler   *handlers.SyncHandler
	ClientHandler *handlers.ClientHandler
	TubeHandler   *handlers.TubeHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	//if err != nil {
	//	return nil, err
	//}

	//dbM := client.Database(root.Mongo.Database)

	cassandra, err := Db(&root)
	if err != nil {
		return nil, err
	}

	dbS, err := sql.Open(root.Postgre.Driver, root.Postgre.DataSourceName)
	if err != nil {
		return nil, err
	}

	//mongoChecker := mgo.NewHealthChecker(dbM)
	sqlChecker := s.NewHealthChecker(dbS)
	healthHandler := health.NewHandler(sqlChecker)

	tubeCategory := tube_category_service.NewCategoryTubeService(root.Key)

	tubeService := tube_service.NewTubeService(root.Key)
	tubeHandler := handlers.NewTubeHandler(tubeService)

	//channelCollectionName := "channel"
	//channelSyncCollectionName := "channelSync"
	//playlistCollectionName := "playlist"
	//playlistVideoCollectionName := "playlistVideo"
	//videoCollectionName := "video"
	//categoryCollectionName := "category"
	//repo := sync_repository.NewMongoVideoRepository(db, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName)
	//syncService := sync_service.NewDefaultSyncService(tubeService, repo)

	repo := sync_repository.NewCassandraVideoRepository(cassandra)
	syncService := sync_service.NewDefaultCassandraSyncService(tubeService, repo)

	syncHandler := handlers.NewSyncHandler(syncService)

	//clientService := client_repository.NewMongoVideoService(db, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName, *tubeCategory)
	clientService := client_repository.NewCassandraVideoRepository(cassandra, *tubeCategory)
	clientHandler := handlers.NewClientHandler(clientService)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		ClientHandler: clientHandler,
		SyncHandler:   syncHandler,
		TubeHandler:   tubeHandler,
	}, nil
}
