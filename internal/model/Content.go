package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CatFact struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"64b1cddaf42a0b75a63f83a9"`
	Fact   string             `json:"fact" bson:"fact" example:"Cats can rotate their ears 180 degrees."`
	Length int                `json:"length" bson:"length" example:"47"`
}
