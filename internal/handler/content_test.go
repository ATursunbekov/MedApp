package handler

import (
	"errors"
	"github.com/ATursunbekov/MedApp/internal/service"
	"github.com/ATursunbekov/MedApp/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := &Handler{} // no service needed

	t.Run("returns ID and status from context", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/content/check", nil)
		w := httptest.NewRecorder()

		r := gin.Default()
		r.GET("/content/check", func(c *gin.Context) {
			// Simulate middleware
			c.Set(savedID, "user123")
			c.Set(status, true)
			h.login(c)
		})

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "user123")
		assert.Contains(t, w.Body.String(), "true")
	})
}

func TestHandler_getCatFacts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContent := mocks.NewMockContent(ctrl)
	service := service.Service{Content: mockContent}
	h := &Handler{service: &service}

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			mockBehavior: func() {
				mockContent.EXPECT().
					GetCatFacts().
					Return([]bson.M{{"text": "Cats purr to communicate"}}, nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "Cats purr",
		},
		{
			name: "Service error",
			mockBehavior: func() {
				mockContent.EXPECT().
					GetCatFacts().
					Return(nil, errors.New("service error"))
			},
			expectedCode:   http.StatusInternalServerError,
			expectedSubstr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/getCatFacts", nil)
			w := httptest.NewRecorder()
			r := gin.Default()
			r.GET("/getCatFacts", h.getCatFacts)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}
