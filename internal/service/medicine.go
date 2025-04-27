package service

import (
	"MedApp/internal/model"
	"MedApp/internal/repository"
	"context"
)

type MedicineService struct {
	repo repository.Medicine
}

func NewMedicineService(repo repository.Medicine) *MedicineService {
	return &MedicineService{repo: repo}
}

func (s *MedicineService) Create(ctx context.Context, medicine *model.Medicine) (string, error) {
	id, err := s.repo.Create(ctx, medicine)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *MedicineService) GetByID(ctx context.Context, id string) (*model.Medicine, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MedicineService) GetAll(ctx context.Context) ([]*model.Medicine, error) {
	return s.repo.GetAll(ctx)
}

func (s *MedicineService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
