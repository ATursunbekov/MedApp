package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
	ID           primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string              `json:name bson:"name"`
	Age          int                 `json:age bson:"age"`
	Sex          string              `json:sex bson:"sex"`
	Speciality   string              `json:speciality bson:"speciality"`
	Phone        string              `json:phone bson:"phone"`
	Email        string              `json:email bson:"email"`
	Password     string              `json:password bson:"password"`
	WeekSchedule []WeekScheduleModel `json:weekSchedule" bson:"weekSchedule"`
}

type WeekScheduleModel struct {
	Date  string   `json:date bson:"date"`
	Slots []string `json:slots bson:"slots"`
}

type DoctorInput struct {
	Email    string `json:email bson:"email"`
	Password string `json:password bson:"password"`
}
