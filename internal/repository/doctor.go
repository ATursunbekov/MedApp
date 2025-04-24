package repository

import (
	"context"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DoctorRepository struct {
	doctorDB *mongo.Collection
}

func NewDoctorRepository(db *mongo.Database) *DoctorRepository {
	return &DoctorRepository{
		doctorDB: db.Collection("doctors"),
	}
}

func (repo *DoctorRepository) GetAllDoctors() ([]model.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := repo.doctorDB.Find(ctx, bson.M{})
	if err != nil {
		logrus.Errorf("Error getting all doctors: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var doctors []model.Doctor

	for cursor.Next(ctx) {
		var doc model.Doctor
		if err := cursor.Decode(&doc); err != nil {
			logrus.Errorf("Error getting all doctors: %v", err)
			return nil, err
		}
		doctors = append(doctors, doc)
	}
	if err := cursor.Err(); err != nil {
		logrus.Errorf("Error getting all doctors: %v", err)
		return nil, err
	}
	return doctors, nil
}
