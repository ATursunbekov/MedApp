package service

import (
	"context"
	"encoding/json"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/repository"
	redisdb "github.com/ATursunbekov/MedApp/pkg/redis"
	"time"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) FindClientByID(id string) (*model.Client, error) {
	cacheKey := "client_profile:" + id

	cached, err := redisdb.Client.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var client model.Client
		if err := json.Unmarshal([]byte(cached), &client); err == nil {
			return &client, nil
		}
	}

	clientDB, err := s.repo.FindClientByID(id)
	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(clientDB)
	redisdb.Client.Set(context.Background(), cacheKey, bytes, 10*time.Minute)
	return clientDB, nil
}

func (s *ProfileService) FindDoctorByID(id string) (*model.Doctor, error) {
	cacheKey := "doctor_profile:" + id

	cached, err := redisdb.Client.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var doctor model.Doctor
		if err := json.Unmarshal([]byte(cached), &doctor); err == nil {
			return &doctor, nil
		}
	}

	doctorDB, err := s.repo.FindDoctorByID(id)
	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(doctorDB)
	redisdb.Client.Set(context.Background(), cacheKey, bytes, 10*time.Minute)
	return doctorDB, nil
}
