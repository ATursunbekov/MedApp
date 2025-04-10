package service

import (
	"MedApp/internal/model"
	"MedApp/internal/repository"
	"encoding/json"
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
	return h.repo.GetCatFact(id)
}
