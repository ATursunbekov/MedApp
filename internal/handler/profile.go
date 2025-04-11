package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getClientProfile(c *gin.Context) {
	var body map[string]interface{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	userID, ok := body["id"].(string)
	if !ok || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'id'"})
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

	c.JSON(http.StatusOK, client)
}

func (h *Handler) getDoctorProfile(c *gin.Context) {
	var body map[string]interface{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	userID, ok := body["id"].(string)
	if !ok || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'id'"})
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

	c.JSON(http.StatusOK, doctor)
}
