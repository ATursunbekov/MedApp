package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingModel struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClientID string             `json:"client_id" bson:"client_id"`
	DoctorID string             `json:"doctor_id" bson:"doctor_id"`
	Date     string             `json:"date" bson:"date"`
	Time     string             `json:"time" bson:"time"`
	Status   string             `json:"status" bson:"status"` //booked / canceled
}

type BookingInput struct {
	DoctorID string `json:"doctor_id" bson:"doctor_id"`
	Date     string `json:"date" bson:"date"`
	Time     string `json:"time" bson:"time"`
}
