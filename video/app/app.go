package app

import (
	"context"
	"database/sql"
	"github.com/core-go/health"
	cas "github.com/core-go/health/cassandra"
	mgo "github.com/core-go/health/mongo"
	s "github.com/core-go/health/sql"
	_ "github.com/lib/pq"
	"go-service/video/client_cassandra"
	mg "go-service/video/client_mongo"
	"go-service/video/sync"
	"go-service/video/sync_cassandra"
	"go-service/video/sync_mongo"
	"go-service/video/sync_postgre"
	"go-service/video/youtube"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/video/handlers"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	SyncHandler   *handlers.SyncHandler
	ClientHandler *handlers.ClientHandler
	TubeHandler   *handlers.TubeHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	var healthHandler *health.Handler
	var clientHandler *handlers.ClientHandler

	var syncHandler *handlers.SyncHandler

	tubeCategory := youtube.NewCategoryTubeService(root.Key)

	tubeService := youtube.NewTubeService(root.Key)
	tubeHandler := handlers.NewTubeHandler(tubeService)

	switch root.OpenDb {
	case 1:
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
		if err != nil {
			return nil, err
		}

		mongoDb := client.Database(root.Mongo.Database)
		mongoChecker := mgo.NewHealthChecker(mongoDb)
		healthHandler = health.NewHandler(mongoChecker)
		channelCollectionName := "channel"
		channelSyncCollectionName := "channelSync"
		playlistCollectionName := "playlist"
		playlistVideoCollectionName := "playlistVideo"
		videoCollectionName := "video"
		categoryCollectionName := "category"
		repo := sync_mongo.NewMongoVideoRepository(mongoDb, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = handlers.NewSyncHandler(syncService)
		clientService := mg.NewMongoVideoService(mongoDb, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName, *tubeCategory)
		clientHandler = handlers.NewClientHandler(clientService)
		break
	case 2:
		cassDb, err := Db(&root)
		if err != nil {
			return nil, err
		}
		casChecker := cas.NewHealthChecker(cassDb)
		healthHandler = health.NewHandler(casChecker)
		repo := sync_cassandra.NewCassandraVideoRepository(cassDb)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = handlers.NewSyncHandler(syncService)
		clientService := client_cassandra.NewCassandraVideoService(cassDb, *tubeCategory)
		clientHandler = handlers.NewClientHandler(clientService)
		break
	case 3:
		postgreDB, err := sql.Open(root.Postgre.Driver, root.Postgre.DataSourceName)
		if err != nil {
			return nil, err
		}
		sqlChecker := s.NewHealthChecker(postgreDB)
		healthHandler = health.NewHandler(sqlChecker)
		repo := sync_postgre.NewSQLVideoRepository(postgreDB)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = handlers.NewSyncHandler(syncService)
		//clientService := client_postgre.NewPostgreVideoService(postgreDB, *tubeCategory)
		//clientHandler = handlers.NewClientHandler(clientService)
		break
	default:
		break
	}

	return &ApplicationContext{
		HealthHandler: healthHandler,
		ClientHandler: clientHandler,
		SyncHandler:   syncHandler,
		TubeHandler:   tubeHandler,
	}, nil
}
