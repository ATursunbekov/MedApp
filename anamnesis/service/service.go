package service

import (
	"github.com/ATursunbekov/MedApp/anamnesis/models"
	"github.com/ATursunbekov/MedApp/anamnesis/repository"
)

type AnamnesisServer struct {
	repo repository.AnamnesisRepository
}

func NewAnamnesisServer(repo repository.AnamnesisRepository) *AnamnesisServer {
	return &AnamnesisServer{repo: repo}
}

func (s *AnamnesisServer) SaveSession(anamnesis models.Anamnesis) (string, error) {
	return s.repo.SaveSession(anamnesis)
}
