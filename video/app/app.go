package app

import (
	"context"
	"go-service/video/cassandra"
	"go-service/video/sync"
	"go-service/video/sync_cassandra"
	//"database/sql"

	"go-service/video/youtube"

	"github.com/core-go/health"
	//cas "github.com/core-go/health/cassandra"
	mgo "github.com/core-go/health/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go-service/video/cassandra"
	//"go-service/video/sync"
	//"go-service/video/sync_cassandra"
	//s "github.com/core-go/health/sql"
	_ "github.com/lib/pq"

	"go-service/video/handlers"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	SyncHandler   *handlers.SyncHandler
	ClientHandler *handlers.ClientHandler
	TubeHandler   *handlers.TubeHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	if err != nil {
		return nil, err
	}

	dbM := client.Database(root.Mongo.Database)
	mongoChecker := mgo.NewHealthChecker(dbM)

	cassDb, err := Db(&root)
	if err != nil {
		return nil, err
	}

	//dbS, err := sql.Open(root.Postgre.Driver, root.Postgre.DataSourceName)
	//if err != nil {
	//	return nil, err
	//}
	//sqlChecker := s.NewHealthChecker(dbS)

	healthHandler := health.NewHandler(mongoChecker)

	tubeCategory := youtube.NewCategoryTubeService(root.Key)

	tubeService := youtube.NewTubeService(root.Key)
	tubeHandler := handlers.NewTubeHandler(tubeService)

	//channelCollectionName := "channel"
	//channelSyncCollectionName := "channelSync"
	//playlistCollectionName := "playlist"
	//playlistVideoCollectionName := "playlistVideo"
	//videoCollectionName := "video"
	//categoryCollectionName := "category"
	//repo := sync_mongo.NewMongoVideoRepository(dbM, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName)
	//syncService := sync.NewDefaultSyncService(tubeService, repo)
	//clientService := mg.NewMongoVideoService(dbM, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName, *tubeCategory)

	repo := sync_cassandra.NewCassandraVideoRepository(cassDb)
	syncService := sync.NewDefaultSyncService(tubeService, repo)
	clientService := cassandra.NewCassandraVideoService(cassDb, *tubeCategory)

	//repo := sync_repository.NewSQLVideoRepository(dbS)
	//syncService := sync_service.NewDefaultSQLSyncService(tubeService, repo)

	syncHandler := handlers.NewSyncHandler(syncService)

	clientHandler := handlers.NewClientHandler(clientService)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		ClientHandler: clientHandler,
		SyncHandler:   syncHandler,
		TubeHandler:   tubeHandler,
	}, nil
}
