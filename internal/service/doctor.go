package service

import (
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/repository"
	"github.com/sirupsen/logrus"
)

var availableSlots = []string{
	"09:00", "10:00", "11:00", "12:00", "13:00",
	"14:00", "15:00", "16:00", "17:00", "18:00",
}

type DoctorService struct {
	repo repository.Repository
}

func NewDoctorService(repo repository.Repository) *DoctorService {
	return &DoctorService{
		repo: repo,
	}
}

func (s *DoctorService) GetAllDoctors() ([]model.Doctor, error) {
	return s.repo.GetAllDoctors()
}

func (s *DoctorService) GetDoctorFreeSlots(id string, date string) ([]string, []string, error) {
	doctor, err := s.repo.FindDoctorByID(id)
	if err != nil {
		logrus.Errorf("failed to find doctor: %v", err)
		return []string{}, []string{}, err
	}

	if doctor.WeekSchedule == nil {
		return []string{}, availableSlots, nil
	}

	dateSlots, check := FindScheduleByDate(doctor.WeekSchedule, date)
	if !check {
		return []string{}, availableSlots, nil
	}

	freeTime := getFreeSlots(availableSlots, dateSlots.Slots)

	return dateSlots.Slots, freeTime, nil
}

// TODO: Filtration logic

func FindScheduleByDate(schedules []model.WeekScheduleModel, date string) (*model.WeekScheduleModel, bool) {
	for _, s := range schedules {
		if s.Date == date {
			return &s, true // Found
		}
	}
	return nil, false // Not found
}

func getFreeSlots(allSlots, occupied []string) []string {
	// Convert occupied slice to a map for O(1) lookups
	occupiedMap := make(map[string]bool)
	for _, slot := range occupied {
		occupiedMap[slot] = true
	}

	// Filter out the occupied slots
	var freeSlots []string
	for _, slot := range allSlots {
		if !occupiedMap[slot] {
			freeSlots = append(freeSlots, slot)
		}
	}
	return freeSlots
}
