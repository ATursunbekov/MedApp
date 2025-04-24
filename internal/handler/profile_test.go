package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/service"
	"github.com/ATursunbekov/MedApp/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getDoctorProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProfile(ctrl)
	mainService := service.Service{
		Profile: mockService,
	}

	h := Handler{
		service: &mainService,
	}

	objID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tests := []struct {
		name           string            // Name of the sub-test
		body           map[string]string // Request body to send
		mockBehavior   func()            // Behavior to mock on mockServices
		expectedCode   int               // Expected HTTP status code
		expectedSubstr string            // Substring expected to be in response body
	}{
		{
			name: "Success",
			body: map[string]string{"id": "doc123"},
			mockBehavior: func() {
				mockService.EXPECT().FindDoctorByID("doc123").Return(&model.Doctor{ID: objID, Name: "Dr. House"}, nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "Dr. House",
		},
		{
			name:           "Missing ID",
			body:           map[string]string{},
			mockBehavior:   func() {}, // No call expected to service layer
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "Missing or invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			bodyBytes, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, "/content/doctor/profile", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/content/doctor/profile", h.getDoctorProfile)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}

func TestHandler_getClientProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProfile(ctrl)
	mainService := service.Service{
		Profile: mockService,
	}
	h := Handler{
		service: &mainService,
	}

	objID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tests := []struct {
		name           string
		body           map[string]string
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			body: map[string]string{"id": "doc123"},
			mockBehavior: func() {
				mockService.EXPECT().FindClientByID("doc123").Return(&model.Client{ID: objID, Name: "Dr. House"}, nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "Dr. House",
		},
		{
			name:           "Missing ID",
			body:           map[string]string{},
			mockBehavior:   func() {},
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "Missing or invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			bodyBytes, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, "/content/client/profile", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r := gin.Default()
			r.POST("/content/client/profile", h.getClientProfile)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}
