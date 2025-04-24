package handler

import (
	"github.com/ATursunbekov/MedApp/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// User
// @Summary      Login Client
// @Tags         AuthClient
// @Description  Authenticates a client and returns a JWT token
// @Accept       json
// @Produce      json
// @Param        input  body      model.ClientInput  true  "Client login credentials"
// @Success      200    {object}  map[string]string  "token response"
// @Failure      400    {object}  map[string]string  "error response"
// @Router       /auth/client/login [post]
func (h *Handler) loginClient(c *gin.Context) {
	var input model.ClientInput
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.LoginClient(input)
	if err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary      Register Client
// @Description  Registers a new client and returns an access token
// @Tags         AuthClient
// @Accept       json
// @Produce      json
// @Param        input  body      model.Client  true  "Client registration details"
// @Success      200    {object}  map[string]string  "access token"
// @Failure      400    {object}  map[string]string  "bad request"
// @Failure      500    {object}  map[string]string  "internal error"
// @Router       /auth/client/register [post]
func (h *Handler) registerClient(c *gin.Context) {
	var input model.Client
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

// @Summary      Login Doctor
// @Tags         AuthDoctor
// @Description  Authenticates a doctor and returns a JWT token
// @Accept       json
// @Produce      json
// @Param        input  body      model.DoctorInput  true  "Doctor login credentials"
// @Success      200    {object}  map[string]string  "token response"
// @Failure      400    {object}  map[string]string  "error response"
// @Router       /auth/doctor/login [post]
// Doctor
func (h *Handler) loginDoctor(c *gin.Context) {
	var input model.DoctorInput
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

// @Summary      Register Doctor
// @Description  Registers a new doctor and returns an access token
// @Tags         AuthDoctor
// @Accept       json
// @Produce      json
// @Param        input  body      model.Doctor  true  "Doctor registration details"
// @Success      200    {object}  map[string]string  "access token"
// @Failure      400    {object}  map[string]string  "bad request"
// @Failure      500    {object}  map[string]string  "internal error"
// @Router       /auth/doctor/register [post]
func (h *Handler) registerDoctor(c *gin.Context) {
	var input model.Doctor
	if err := c.ShouldBind(&input); err != nil {
		logrus.Errorf("Login Client Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
