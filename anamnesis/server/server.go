package server

import (
	"context"
	"github.com/ATursunbekov/MedApp/anamnesis/models"
	"github.com/ATursunbekov/MedApp/anamnesis/repository"
	service2 "github.com/ATursunbekov/MedApp/anamnesis/service"
	pb "github.com/ATursunbekov/MedApp/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnamnesisServer struct {
	pb.UnimplementedSessionServiceServer
	repo repository.AnamnesisRepository
}

func NewAnamnesisServer(db *mongo.Database) AnamnesisServer {
	return AnamnesisServer{
		repo: *repository.NewAnamnesisRepository(db),
	}
}

func (s *AnamnesisServer) SaveSession(ctx context.Context, req *pb.SaveSessionRequest) (*pb.SaveSessionResponse, error) {
	input := models.Anamnesis{
		Date:   req.Timestamp,
		UserID: req.UserId,
		Notes:  req.Notes,
	}

	service := service2.NewAnamnesisServer(s.repo)
	res, err := service.SaveSession(input)
	if err != nil {
		logrus.Errorf("Error saving session: %v", err)
		return nil, err
	}

	return &pb.SaveSessionResponse{
		Status: res,
	}, nil
}

func (s *AnamnesisServer) mustEmbedUnimplementedChatServiceServer() {}
