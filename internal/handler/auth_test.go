package handler

import (
	"bytes"
	"encoding/json"
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/ATursunbekov/MedApp/internal/service"
	"github.com/ATursunbekov/MedApp/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_loginClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)

	service := service.Service{
		Authorization: mockAuth,
	}

	h := &Handler{service: &service}

	tests := []struct {
		name           string
		input          interface{}
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			input: model.ClientInput{
				Email:    "test@gmail.com",
				Password: "12345",
			},
			mockBehavior: func() {
				mockAuth.EXPECT().
					LoginClient(model.ClientInput{Email: "test@gmail.com", Password: "12345"}).
					Return("test-token", nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "test-token",
		},
		{
			name:           "Invalid JSON",
			input:          `invalid-json`,
			mockBehavior:   func() {},
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			var body []byte
			switch v := tt.input.(type) {
			case string:
				body = []byte(v)
			default:
				body, _ = json.Marshal(v)
			}

			req, _ := http.NewRequest(http.MethodPost, "/auth/client/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r := gin.Default()
			r.POST("/auth/client/login", h.loginClient)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}

func TestHandler_registerClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)
	service := service.Service{
		Authorization: mockAuth,
	}
	h := &Handler{service: &service}

	tests := []struct {
		name           string
		input          interface{}
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			input: model.Client{
				Email:    "test@gmail.com",
				Password: "12345",
				Name:     "John",
				Age:      30,
				Sex:      "male",
				Phone:    "13812345678",
			},
			mockBehavior: func() {
				mockAuth.EXPECT().
					CreateClient(model.Client{Email: "test@gmail.com",
						Password: "12345",
						Name:     "John",
						Age:      30,
						Sex:      "male",
						Phone:    "13812345678"}).
					Return("access-token", nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "access-token",
		},
		{
			name:           "Invalid JSON",
			input:          `invalid-json`,
			mockBehavior:   func() {},
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			var body []byte
			switch v := tt.input.(type) {
			case string:
				body = []byte(v)
			default:
				body, _ = json.Marshal(v)
			}

			req, _ := http.NewRequest(http.MethodPost, "/auth/client/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r := gin.Default()
			r.POST("/auth/client/register", h.registerClient)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}

func TestHandler_loginDoctor(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)

	service := service.Service{
		Authorization: mockAuth,
	}
	h := &Handler{service: &service}

	tests := []struct {
		name           string
		input          interface{}
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			input: model.DoctorInput{
				Email:    "doc@example.com",
				Password: "12345",
			},
			mockBehavior: func() {
				mockAuth.EXPECT().
					LoginDoctor(model.DoctorInput{Email: "doc@example.com", Password: "12345"}).
					Return("doctor-token", nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "doctor-token",
		},
		{
			name:           "Invalid JSON",
			input:          `bad-json-format`,
			mockBehavior:   func() {},
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			var body []byte
			switch v := tt.input.(type) {
			case string:
				body = []byte(v)
			default:
				body, _ = json.Marshal(v)
			}

			req, _ := http.NewRequest(http.MethodPost, "/auth/doctor/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r := gin.Default()
			r.POST("/auth/doctor/login", h.loginDoctor)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}

func TestHandler_registerDoctor(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthorization(ctrl)
	service := service.Service{
		Authorization: mockAuth,
	}
	h := &Handler{service: &service}

	tests := []struct {
		name           string
		input          interface{}
		mockBehavior   func()
		expectedCode   int
		expectedSubstr string
	}{
		{
			name: "Success",
			input: model.Doctor{
				Name:       "Dr. John",
				Email:      "drjohn@example.com",
				Password:   "docpass",
				Sex:        "male",
				Age:        40,
				Phone:      "777123456",
				Speciality: "okulist",
			},
			mockBehavior: func() {
				mockAuth.EXPECT().
					CreateDoctor(model.Doctor{
						Name:       "Dr. John",
						Email:      "drjohn@example.com",
						Password:   "docpass",
						Sex:        "male",
						Age:        40,
						Phone:      "777123456",
						Speciality: "okulist",
					}).
					Return("access-token", nil)
			},
			expectedCode:   http.StatusOK,
			expectedSubstr: "access-token",
		},
		{
			name:           "Invalid JSON",
			input:          `{{invalid_payload`,
			mockBehavior:   func() {},
			expectedCode:   http.StatusBadRequest,
			expectedSubstr: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			var body []byte
			switch v := tt.input.(type) {
			case string:
				body = []byte(v)
			default:
				body, _ = json.Marshal(v)
			}

			req, _ := http.NewRequest(http.MethodPost, "/auth/doctor/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r := gin.Default()
			r.POST("/auth/doctor/register", h.registerDoctor)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedSubstr)
		})
	}
}
