package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingModel struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"64a9b8b7e6f5d9d5a5c6b2f0"`
	ClientID string             `json:"client_id" bson:"client_id" example:"643ebec937a2d9b4b2645f3c"`
	DoctorID string             `json:"doctor_id" bson:"doctor_id" example:"643ebec937a2d9b4b2645f3d"`
	Date     string             `json:"date" bson:"date" example:"12-04-2025"`
	Time     string             `json:"time" bson:"time" example:"15:30"`
	Status   string             `json:"status" bson:"status" example:"booked"` // Options: booked / canceled
}

type BookingInput struct {
	DoctorID string `json:"doctor_id" bson:"doctor_id" example:"643ebec937a2d9b4b2645f3d"`
	Date     string `json:"date" bson:"date" example:"12-04-2025"`
	Time     string `json:"time" bson:"time" example:"15:30"`
}
