package handler

import (
	"MedApp/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	// test endpoint
	router.GET("/getCatFacts", h.getCatFacts)
	router.GET("/getCatFact/:id", h.getCatFact)

	// Auth group
	auth := router.Group("/auth")
	{
		client := auth.Group("/client")
		{
			client.POST("/register", h.registerClient)
			client.POST("/login", h.loginClient)
		}

		doctor := auth.Group("/doctor")
		{
			doctor.POST("/register", h.registerDoctor)
			doctor.POST("/login", h.loginDoctor)
		}
	}

	content := router.Group("/content", h.JWTMiddleware)
	{
		content.GET("/check", h.login)

		client := content.Group("/client")
		{
			client.GET("/profile", h.getClientProfile)
			client.POST("/book", h.bookTimeSlot)
		}

		doctor := content.Group("/doctor")
		{
			doctor.GET("profile", h.getDoctorProfile)
		}

		//TODO: General apis
		content.GET("/getDoctors", h.getAllDoctors)
		content.GET("/getSchedule", h.getFreeTimeSlots)
	}
	// For saving facts in db
	//h.service.SaveCatFacts()

	return router
}
