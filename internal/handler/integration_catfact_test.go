package handler

import (
	"context"
	"encoding/json"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/repository"
	"github.com/ATursunbekov/MedApp/internal/service"
	"github.com/ATursunbekov/MedApp/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) *repository.Repository {
	cfg := repository.MongoConfig{
		URI:     "mongodb://admin:secret@localhost:27017", // Make sure Redis & Mongo are running
		DBName:  "medapp_test",
		Timeout: 10 * time.Second,
	}

	db, err := repository.ConnectMongo(cfg)
	if err != nil {
		t.Fatalf("❌ Failed to connect to test MongoDB: %v", err)
	}

	// Drop old test data
	if err := db.Collection("catfacts").Drop(context.Background()); err != nil {
		t.Fatalf("❌ Failed to clear catfacts: %v", err)
	}

	return repository.NewRepository(db)
}

func TestIntegration_getCatFact(t *testing.T) {
	// 🧪 Step 1: Disable Redis logic in service layer (e.g., skip cache hit logic)
	os.Setenv("TEST_ENV", "true")

	// 🔁 Optional: if you want Redis working, initialize Redis connection (safe even if unused)
	redis.InitRedis()

	// 🔇 Step 2: Set Gin to test mode to suppress default output
	gin.SetMode(gin.TestMode)

	// 📦 Step 3: Setup real MongoDB and clean test DB
	repo := setupTestDB(t)

	// 🧪 Step 4: Create test cat fact data and insert into Mongo
	objID := primitive.NewObjectID() // generate random ObjectID
	expected := model.CatFact{
		ID:   objID,
		Fact: "Cats sleep for 70% of their lives",
	}

	// 🗃 Insert directly into Mongo (or use repo.SaveCatFact logic if you prefer abstraction)
	err := repo.Content.SaveCatFact(bson.M{
		"_id":  expected.ID,
		"Fact": expected.Fact,
	})
	assert.NoError(t, err)

	// 🧠 Step 5: Wire up actual app logic: repo → service → handler
	svc := service.NewService(repo)
	h := NewHandler(svc)

	// 🌐 Step 6: Create a test router and register the real endpoint
	r := gin.Default()
	r.GET("/catfacts/:id", h.getCatFact)

	// 🌍 Step 7: Simulate a real HTTP request to that route
	req, _ := http.NewRequest("GET", "/catfacts/"+objID.Hex(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// ✅ Step 8: Verify HTTP response
	assert.Equal(t, http.StatusOK, w.Code) // check for HTTP 200

	// 📥 Step 9: Parse JSON response body into a Go struct
	var actual model.CatFact
	err = json.NewDecoder(w.Body).Decode(&actual)
	assert.NoError(t, err)

	// 🔍 Step 10: Validate the content of the response
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Fact, actual.Fact)
}
