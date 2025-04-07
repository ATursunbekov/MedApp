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
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "SUCCESS",
		})
	})

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
	}

	return router
}

func (h *Handler) login(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)
	c.JSON(200, gin.H{
		"ID":     userID,
		"status": isClient,
	})
}
