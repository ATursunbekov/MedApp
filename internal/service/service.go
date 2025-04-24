package service

import (
	"github.com/ATursunbekov/MedApp/internal/model"
	repository "github.com/ATursunbekov/MedApp/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Authorization interface {
	CreateClient(client model.Client) (string, error)
	CreateDoctor(doctor model.Doctor) (string, error)
	LoginClient(input model.ClientInput) (string, error)
	LoginDoctor(input model.DoctorInput) (string, error)
}

type Profile interface {
	FindClientByID(id string) (*model.Client, error)
	FindDoctorByID(id string) (*model.Doctor, error)
}

type Doctor interface {
	GetAllDoctors() ([]model.Doctor, error)
	GetDoctorFreeSlots(id string, date string) ([]string, []string, error)
}

type Booking interface {
	BookSession(booking model.BookingModel) error
}

type Content interface {
	GetCatFacts() ([]bson.M, error)
	SaveCatFacts() error
	GetCatFact(id string) (*model.CatFact, error)
}

type Service struct {
	Authorization
	Profile
	Doctor
	Booking
	Content
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Profile:       NewProfileService(repo),
		Doctor:        NewDoctorService(*repo),
		Booking:       NewBookingService(repo),
		Content:       NewContent(*repo),
	}
}

type AllServices interface {
	Authorization
	Profile
	Doctor
	Booking
	Content
}
