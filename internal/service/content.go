package service

import (
	"context"
	"encoding/json"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/repository"
	redisdb "github.com/ATursunbekov/MedApp/pkg/redis"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type ContentService struct {
	repo repository.Repository
}

func NewContent(repo repository.Repository) *ContentService {
	return &ContentService{
		repo: repo,
	}
}

func (h *ContentService) GetCatFacts() ([]bson.M, error) {
	return h.repo.GetAllCatFacts()
}

func (h *ContentService) SaveCatFacts() error {
	ticker := time.NewTicker(5 * time.Second)

	for {
		response, err := http.Get("http://catfact.ninja/fact")
		if err != nil {
			return err
		}
		var catFact bson.M
		if err := json.NewDecoder(response.Body).Decode(&catFact); err != nil {
			return err
		}

		err = h.repo.SaveCatFact(catFact)
		if err != nil {
			return err
		}
		<-ticker.C
	}
}

func (h *ContentService) GetCatFact(id string) (*model.CatFact, error) {
	cacheKey := "cat_fact:" + id

	cached, err := redisdb.Client.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var fact model.CatFact
		if err := json.Unmarshal([]byte(cached), &fact); err == nil {
			return &fact, nil
		}
	}

	catDB, err := h.repo.GetCatFact(id)
	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(catDB)
	redisdb.Client.Set(context.Background(), cacheKey, bytes, 10*time.Minute)
	return catDB, nil
}
