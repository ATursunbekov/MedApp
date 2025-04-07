package service

import (
	"MedApp/internal/model"
	repository "MedApp/internal/repository"
)

type Authorization interface {
	CreateClient(client model.Client) (string, error)
	CreateDoctor(doctor model.Doctor) (string, error)
	LoginClient(input model.ClientInput) (string, error)
	LoginDoctor(input model.DoctorInput) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
