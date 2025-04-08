package service

import (
	"MedApp/internal/model"
	"MedApp/internal/repository"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(repo repository.Booking) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (s *BookingService) BookSession(booking model.BookingModel) error {
	err := s.repo.BookClientToDoctor(booking)
	if err != nil {
		return err
	}
	return nil
}
