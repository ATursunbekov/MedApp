package repository

import (
	"context"
	"github.com/ATursunbekov/MedApp/anamnesis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AnamnesisRepository struct {
	clientDB *mongo.Collection
	doctorDB *mongo.Collection
}

func NewAnamnesisRepository(db *mongo.Database) *AnamnesisRepository {
	return &AnamnesisRepository{
		clientDB: db.Collection("anamnesises"),
	}
}

func (r *AnamnesisRepository) SaveSession(session models.Anamnesis) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if session.ID.IsZero() {
		session.ID = primitive.NewObjectID()
	}

	res, err := r.clientDB.InsertOne(ctx, session)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AnamnesisRepository) ClientsAllSessions(clientID string) ([]models.Anamnesis, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": clientID}

	cursor, err := r.clientDB.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []models.Anamnesis
	if err = cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}
