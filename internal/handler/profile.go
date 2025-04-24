package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary      Get Client Profile
// @Description  Returns profile information of a client by their ID
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        input  body      map[string]string  true  "JSON with client ID"
// @Success      200    {object}  model.Client
// @Failure      400    {object}  map[string]string  "Missing or invalid ID"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /content/client/profile [post]
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

// @Summary      Get Doctor Profile
// @Description  Returns profile information of a doctor by their ID
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        input  body      map[string]string  true  "JSON with doctor ID"
// @Success      200    {object}  model.Doctor
// @Failure      400    {object}  map[string]string  "Missing or invalid ID"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /content/doctor/profile [post]
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
