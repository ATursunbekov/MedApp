package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI     string
	DBName  string
	Timeout time.Duration
}

var (
	mongoClient *mongo.Client
	once        sync.Once
)

// ConnectMongo establishes a new MongoDB connection
func ConnectMongo(cfg MongoConfig) (*mongo.Database, error) {
	var err error

	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		logrus.Info(cfg.URI)

		clientOptions := options.Client().ApplyURI(cfg.URI)

		mongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Printf("❌ Failed to connect to MongoDB: %v", err)
			return
		}

		// Verify connection
		if err = mongoClient.Ping(ctx, nil); err != nil {
			log.Printf("❌ MongoDB ping failed: %v", err)
			return
		}

		log.Println("✅ Connected to MongoDB")
	})

	if err != nil {
		return nil, err
	}

	return mongoClient.Database(cfg.DBName), nil
}
