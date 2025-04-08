package repository

import (
	"MedApp/internal/model"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookingRepository struct {
	doctorDB  *mongo.Collection
	bookingDB *mongo.Collection
}

func NewBookingRepository(db *mongo.Database) *BookingRepository {
	return &BookingRepository{
		doctorDB:  db.Collection("doctors"),
		bookingDB: db.Collection("bookings"),
	}
}

func (b BookingRepository) BookClientToDoctor(session model.BookingModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	session.ID = primitive.NewObjectID()
	_, err := b.bookingDB.InsertOne(ctx, session)
	if err != nil {
		logrus.Errorf("Error creating session: %v", err)
		return err
	}

	doctorObjectID, err := primitive.ObjectIDFromHex(session.DoctorID)
	if err != nil {
		logrus.Errorf("Error converting doctorID to objectID: %v", err)
		return err
	}

	// Find the doctor by ID
	var doctor model.Doctor
	err = b.doctorDB.FindOne(ctx, bson.M{"_id": doctorObjectID}).Decode(&doctor)
	if err != nil {
		logrus.Errorf("Error finding doctor: %v", err)
		return err
	}

	// Update the doctor's schedule
	found := false
	for i, schedule := range doctor.WeekSchedule {
		if schedule.Date == session.Date {
			//Date found, append time
			doctor.WeekSchedule[i].Slots = append(doctor.WeekSchedule[i].Slots, session.Time)
			found = true
			break
		}
	}

	if !found {
		// Date not found, create new schedule entry
		newDay := model.WeekScheduleModel{
			Date:  session.Date,
			Slots: []string{session.Time},
		}
		doctor.WeekSchedule = append(doctor.WeekSchedule, newDay)
	}

	// Update doctor document in MongoDB
	_, err = b.doctorDB.UpdateOne(
		ctx,
		bson.M{"_id": doctorObjectID},
		bson.M{"$set": bson.M{"weekSchedule": doctor.WeekSchedule}},
	)

	if err != nil {
		logrus.Errorf("Error updating doctor's schedule: %v", err)
		return err
	}

	return nil
}
