package handler

import (
	"MedApp/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// User
func (h *Handler) loginClient(c *gin.Context) {
	var input model.ClientInput
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.service.LoginClient(input)
	if err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) registerClient(c *gin.Context) {
	var input model.Client
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	accessToken, err := h.service.CreateClient(input)
	if err != nil {
		logrus.Errorf("Register Client Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}

// Doctor
func (h *Handler) loginDoctor(c *gin.Context) {
	var input model.DoctorInput
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.service.LoginDoctor(input)
	if err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) registerDoctor(c *gin.Context) {
	var input model.Doctor
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	accessToken, err := h.service.CreateDoctor(input)
	if err != nil {
		logrus.Errorf("Register Doctor Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Login Client Input: %v", input)
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
