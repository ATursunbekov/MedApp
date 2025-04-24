package handler

import (
	"github.com/ATursunbekov/MedApp/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	savedID = "userID"
	status  = "isClient"
)

func (h *Handler) JWTMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		logrus.Infof("Authorization header is not bearer token: %s", authHeader)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userID, isClient, err := service.ParseToken(tokenString)
	if err != nil {
		logrus.Errorf("Error parsing token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.Set(savedID, userID)
	c.Set(status, isClient)

	c.Next()
}

//func (h *Handler) GetProfile(c *gin.Context) {
//	userID := c.GetString("userID")
//	isClient := c.GetBool("isClient")
//
//	if isClient {
//		// fetch client profile
//	} else {
//		// fetch doctor profile
//	}
//}
