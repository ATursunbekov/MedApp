package main

import (
	"MedApp/config"
	"MedApp/internal/handler"
	repository "MedApp/internal/repository"
	service2 "MedApp/internal/service"
	"MedApp/pkg/redis"
	server2 "MedApp/pkg/server"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//fetching env files
	config.LoadEnv()

	//Connecting to db
	db, err := repository.ConnectMongo(repository.MongoConfig{
		URI:     os.Getenv("MONGODB_URI"),
		DBName:  os.Getenv("MONGO_DB_NAME"),
		Timeout: 10 * time.Second,
	})

	//Connecting to Redis
	redis.InitRedis()

	if err != nil {
		logrus.Error("Failed to connect to Mongo: %s", err)
	} else {
		logrus.Info("Connected to Mongo Sucex")
	}

	//project logic setup
	repo := repository.NewRepository(db)
	service := service2.NewService(repo)
	router := handler.NewHandler(service)

	if err := repo.CreateClientEmailIndex(); err != nil {
		logrus.Fatalf("Failed to create unique index on email: %v", err)
	}
	if err := repo.CreateDoctorEmailIndex(); err != nil {
		logrus.Fatalf("Failed to create unique index on email: %v", err)
	}

	server := new(server2.Server)
	if err := server.Run(":8080", router.InitRouter()); err != nil {
		logrus.Error("Failed to start server: %s", err.Error())
	}
}
