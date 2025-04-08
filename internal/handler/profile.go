package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getClientProfile(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)

	if !isClient {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong user status! (Client/Doctor)",
		})
		return
	}

	client, err := h.service.FindClientByID(userID)
	if err != nil {
		logrus.Infof("Error when finding client by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when finding client by ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func (h *Handler) getDoctorProfile(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)

	if isClient {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong user status! (Client/Doctor)",
		})
		return
	}

	doctor, err := h.service.FindDoctorByID(userID)
	if err != nil {
		logrus.Infof("ID: %v", userID)
		logrus.Infof("Error when finding doctor by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when finding doctor by ID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor": doctor,
	})
}
