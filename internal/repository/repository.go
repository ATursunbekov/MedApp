package repository

import (
	"MedApp/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Medicine interface {
	Create(ctx context.Context, medicine *model.Medicine) (string, error)
	GetByID(ctx context.Context, id string) (*model.Medicine, error)
	GetAll(ctx context.Context) ([]*model.Medicine, error)
	Delete(ctx context.Context, id string) error
}

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

type Content interface {
	SaveCatFact(fact bson.M) error
	GetAllCatFacts() ([]bson.M, error)
	GetCatFact(id string) (*model.CatFact, error)
}

type Repository struct {
	Authorization
	Profile
	Doctor
	Booking
	Content
	Medicine
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Profile:       NewProfileRepository(db),
		Doctor:        NewDoctorRepository(db),
		Booking:       NewBookingRepository(db),
		Content:       NewContent(db),
		Medicine:      NewMedicineRepository(db),
	}
}
