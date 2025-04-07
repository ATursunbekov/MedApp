package repository

import (
	"MedApp/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateClient(client model.Client) (string, error)
	CreateDoctor(doctor model.Doctor) (string, error)
	LoginClient(input model.ClientInput) (model.Client, error)
	LoginDoctor(input model.DoctorInput) (model.Doctor, error)
	CreateClientEmailIndex() error
	CreateDoctorEmailIndex() error
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
