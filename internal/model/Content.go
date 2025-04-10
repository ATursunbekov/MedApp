package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CatFact struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fact   string             `json:"fact" bson:"fact"`
	Length int                `json:"length" bson:"length"`
}
