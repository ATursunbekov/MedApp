package main

import (
	"github.com/ATursunbekov/MedApp/anamnesis/server"
	"github.com/ATursunbekov/MedApp/config"
	"github.com/ATursunbekov/MedApp/internal/repository"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/ATursunbekov/MedApp/proto"
	"google.golang.org/grpc"
)

// TODO: Implement gRPC methods here

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	config.LoadEnv()

	db, err := repository.ConnectMongo(repository.MongoConfig{
		URI:     os.Getenv("MONGODB_URI"),
		DBName:  os.Getenv("MONGO_DB_NAME"),
		Timeout: 10 * time.Second,
	})

	s := server.NewAnamnesisServer(db)
	grpcServer := grpc.NewServer()
	pb.RegisterSessionServiceServer(grpcServer, &s)

	log.Println("Anamnesis gRPC Service is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
