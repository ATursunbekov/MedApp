package repository

import (
	"MedApp/internal/model"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AuthRepository struct {
	clientDB *mongo.Collection
	doctorDB *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		clientDB: db.Collection("clients"),
		doctorDB: db.Collection("doctors"),
	}
}

// TODO: Registration
func (a *AuthRepository) CreateClient(client model.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client.ID = primitive.NewObjectID()
	_, err := a.clientDB.InsertOne(ctx, client)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			logrus.Warn("Email already exists")
			return "", fmt.Errorf("email already registered")
		}
		logrus.Errorf("Error creating client: %v", err)
		return "", err
	}

	return client.ID.Hex(), nil
}

func (a *AuthRepository) CreateDoctor(doctor model.Doctor) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doctor.ID = primitive.NewObjectID()
	_, err := a.doctorDB.InsertOne(ctx, doctor)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			logrus.Warn("Email already exists")
			return "", fmt.Errorf("email already registered")
		}
		logrus.Errorf("Error creating client: %v", err)
		return "", err
	}

	return doctor.ID.Hex(), nil
}

// TODO: Login logic

func (a *AuthRepository) LoginClient(input model.ClientInput) (model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var client model.Client

	err := a.clientDB.FindOne(ctx, bson.M{"email": input.Email}).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return client, fmt.Errorf("client does not exist")
		}
		logrus.Errorf("Error getting client: %v", err)
		return client, err
	}
	return client, nil
}

func (a *AuthRepository) LoginDoctor(input model.DoctorInput) (model.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var doctor model.Doctor

	err := a.clientDB.FindOne(ctx, bson.M{"email": input.Email}).Decode(&doctor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doctor, fmt.Errorf("client does not exist")
		}
		logrus.Errorf("Error getting client: %v", err)
		return doctor, err
	}
	return doctor, nil
}

// TODO: Making unique properties
func (a *AuthRepository) CreateClientEmailIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}}, // index on "email" field
		Options: options.Index().SetUnique(true),
	}

	_, err := a.clientDB.Indexes().CreateOne(ctx, indexModel)
	return err
}

func (a *AuthRepository) CreateDoctorEmailIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}}, // index on "email" field
		Options: options.Index().SetUnique(true),
	}

	_, err := a.doctorDB.Indexes().CreateOne(ctx, indexModel)
	return err
}
