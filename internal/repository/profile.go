package repository

import (
	"context"
	"fmt"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProfileRepository struct {
	clientDB *mongo.Collection
	doctorDB *mongo.Collection
}

func NewProfileRepository(db *mongo.Database) *ProfileRepository {
	return &ProfileRepository{
		clientDB: db.Collection("clients"),
		doctorDB: db.Collection("doctors"),
	}
}

func (r *ProfileRepository) FindClientByID(id string) (*model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Warn("Invalid object ID format")
		return nil, fmt.Errorf("invalid ID")
	}

	var client model.Client
	err = r.clientDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Warn("Client not found")
			return nil, fmt.Errorf("does not exist")
		}
		return nil, err
	}

	return &client, nil
}

func (r *ProfileRepository) FindDoctorByID(id string) (*model.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var doctor model.Doctor

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Warn("Invalid object ID format")
		return nil, fmt.Errorf("invalid ID")
	}

	err = r.doctorDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&doctor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Warn("Wrong access token")
			return nil, fmt.Errorf("Doesn't exist")
		}
		return nil, err
	}

	return &doctor, nil
}
