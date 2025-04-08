package service

import (
	"MedApp/internal/model"
	"MedApp/internal/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) FindClientByID(id string) (*model.Client, error) {
	return s.repo.FindClientByID(id)
}

func (s *ProfileService) FindDoctorByID(id string) (*model.Doctor, error) {
	return s.repo.FindDoctorByID(id)
}
