package main

import (
	"github.com/ATursunbekov/MedApp/config"
	"github.com/ATursunbekov/MedApp/internal/handler"
	repository "github.com/ATursunbekov/MedApp/internal/repository"
	service2 "github.com/ATursunbekov/MedApp/internal/service"
	"github.com/ATursunbekov/MedApp/pkg/redis"
	server2 "github.com/ATursunbekov/MedApp/pkg/server"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//swagger link: localhost:8080/swagger/index.html#

// @title           MedApp API
// @version         1.0
// @description     Backend logic for MedApp, main feature is booking to doctor sessions

// @contact.name   Alikhan Tursunbekov
// @contact.email  alikhan.tursunbekov@gmail.com

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
		logrus.Errorf("Failed to connect to Mongo: %s", err)
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
		logrus.Errorf("Failed to start server: %s", err.Error())
	}
}
