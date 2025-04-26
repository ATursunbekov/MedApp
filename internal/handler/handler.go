package handler

import (
	_ "MedApp/docs"
	"MedApp/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			client.POST("/profile", h.getClientProfile)
			client.POST("/book", h.bookTimeSlot)
		}

		doctor := content.Group("/doctor")
		{
			doctor.POST("profile", h.getDoctorProfile)
		}

		// Medicine endpoints
		medicines := content.Group("/medicines")
		{
			medicines.POST("", h.createMedicine)
			medicines.GET("", h.getAllMedicines)
			medicines.GET("/:id", h.getMedicineByID)
			medicines.DELETE("/:id", h.deleteMedicine)
		}

		// General apis
		content.GET("/getDoctors", h.getAllDoctors)
		content.POST("/getSchedule", h.getFreeTimeSlots)
	}

	return router
}
