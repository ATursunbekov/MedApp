package service

import (
	"MedApp/internal/model"
	"MedApp/tests/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMedicineService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	medicine := &model.Medicine{
		ID:          "1",
		Name:        "Paracetamol",
		Description: "Pain reliever",
		Quantity:    100,
	}

	mockRepo.EXPECT().Create(gomock.Any(), medicine).Return(nil)

	err := service.Create(context.Background(), medicine)
	assert.NoError(t, err)
}

func TestMedicineService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	medicine := &model.Medicine{
		ID:          "1",
		Name:        "Paracetamol",
		Description: "Pain reliever",
		Quantity:    100,
	}

	mockRepo.EXPECT().Create(gomock.Any(), medicine).Return(errors.New("database error"))

	err := service.Create(context.Background(), medicine)
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}

func TestMedicineService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	medicine := &model.Medicine{
		ID:          "1",
		Name:        "Paracetamol",
		Description: "Pain reliever",
		Quantity:    100,
	}

	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(medicine, nil)

	result, err := service.GetByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, medicine, result)
}

func TestMedicineService_GetByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	mockRepo.EXPECT().GetByID(gomock.Any(), "1").Return(nil, errors.New("not found"))

	result, err := service.GetByID(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "not found", err.Error())
}

func TestMedicineService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	medicines := []*model.Medicine{
		{
			ID:          "1",
			Name:        "Paracetamol",
			Description: "Pain reliever",
			Quantity:    100,
		},
		{
			ID:          "2",
			Name:        "Ibuprofen",
			Description: "Anti-inflammatory",
			Quantity:    50,
		},
	}

	mockRepo.EXPECT().GetAll(gomock.Any()).Return(medicines, nil)

	result, err := service.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, medicines, result)
}

func TestMedicineService_GetAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	mockRepo.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("database error"))

	result, err := service.GetAll(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
}

func TestMedicineService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	mockRepo.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	err := service.Delete(context.Background(), "1")
	assert.NoError(t, err)
}

func TestMedicineService_Delete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMedicine(ctrl)
	service := NewMedicineService(mockRepo)

	mockRepo.EXPECT().Delete(gomock.Any(), "1").Return(errors.New("not found"))

	err := service.Delete(context.Background(), "1")
	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
}
