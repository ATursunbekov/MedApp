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

type Profile interface {
	FindClientByID(id string) (*model.Client, error)
	FindDoctorByID(id string) (*model.Doctor, error)
}

type Doctor interface {
	GetAllDoctors() ([]model.Doctor, error)
}

type Booking interface {
	BookClientToDoctor(session model.BookingModel) error
}

type Repository struct {
	Authorization
	Profile
	Doctor
	Booking
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Profile:       NewProfileRepository(db),
		Doctor:        NewDoctorRepository(db),
		Booking:       NewBookingRepository(db),
	}
}
