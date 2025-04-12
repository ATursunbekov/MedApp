package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
	ID           primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string              `json:"name" bson:"name" example:"John Doe" binding:"required"`
	Age          int                 `json:"age" bson:"age" example:"22" binding:"required"`
	Sex          string              `json:"sex" bson:"sex" example:"male" binding:"required"`
	Speciality   string              `json:"speciality" bson:"speciality" example:"Okulist" binding:"required"`
	Phone        string              `json:"phone" bson:"phone" example:"+99655742351" binding:"required"`
	Email        string              `json:"email" bson:"email" example:"some@gmail.com" binding:"required"`
	Password     string              `json:"password" bson:"password" example:"*****" binding:"required"`
	WeekSchedule []WeekScheduleModel `json:"weekSchedule" bson:"weekSchedule"`
}

type WeekScheduleModel struct {
	Date  string   `json:"date" bson:"date"`
	Slots []string `json:"slots" bson:"slots"`
}

type DoctorInput struct {
	Email    string `json:"email" example:"user@example.com" bson:"email" binding:"required"`
	Password string `json:"password" example:"securePass123" bson:"password" binding:"required"`
}

// For getting doctor's free slots
type DoctorSchedule struct {
	ID   string `json:"id" bson:"id"`
	Date string `json:"date" bson:"date"` //date format: (00-00-00)
}
