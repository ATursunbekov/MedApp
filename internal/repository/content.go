package repository

import (
	"context"
	"fmt"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ContentRepository struct {
	db *mongo.Database
}

func NewContent(db *mongo.Database) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

func (c *ContentRepository) SaveCatFact(fact bson.M) error {
	collection := c.db.Collection("catFacts")
	_, err := collection.InsertOne(context.TODO(), fact)
	if err != nil {
		return err
	}

	return nil
}

func (c *ContentRepository) GetAllCatFacts() ([]bson.M, error) {
	collection := c.db.Collection("catFacts")

	query := bson.M{}
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}

func (c *ContentRepository) GetCatFact(id string) (*model.CatFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Warn("Invalid object ID format")
		return nil, fmt.Errorf("invalid ID")
	}

	var catFact model.CatFact
	err = c.db.Collection("catFacts").FindOne(ctx, bson.M{"_id": objectID}).Decode(&catFact)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Warn("Client not found")
			return nil, fmt.Errorf("does not exist")
		}
		return nil, err
	}

	return &catFact, nil
}
