package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Anamnesis struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID string             `bson:"user_id"`
	Date   string             `bson:"date"`
	Notes  string             `bson:"notes"`
}
