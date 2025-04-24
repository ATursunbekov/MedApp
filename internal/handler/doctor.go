package handler

import (
	"MedApp/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary      Get All Doctors
// @Description  Returns a list of all registered doctors
// @Tags         ContentActions
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "List of doctors or nil"
// @Failure      500  {object}  map[string]string  "Internal error"
// @Security     ApiKeyAuth
// @Router       /content/getDoctors [get]
func (h *Handler) getAllDoctors(c *gin.Context) {
	doctors, err := h.service.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get all doctors",
		})
		return
	}

	if doctors == nil {
		c.JSON(http.StatusOK, gin.H{
			"doctors": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctors": doctors,
	})
}

// @Summary      Get Free Time Slots
// @Description  Returns free and occupied time slots for a doctor on a given day
// @Tags         ContentActions
// @Accept       json
// @Produce      json
// @Param        input  body      model.DoctorSchedule  true  "Doctor ID and date"
// @Success      200    {object}  map[string]interface{}  "Available and occupied slots"
// @Failure      400    {object}  map[string]string  "Bad request"
// @Failure      500    {object}  map[string]string  "Internal error"
// @Security     ApiKeyAuth
// @Router       /content/getSchedule [post]
func (h *Handler) getFreeTimeSlots(c *gin.Context) {
	var req model.DoctorSchedule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}

	occupied, freeSlots, err := h.service.GetDoctorFreeSlots(req.ID, req.Date)
	if err != nil {
		logrus.Errorf("Get doctor free slot error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get free slots",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"date":         req.Date,
		"occupiedTime": occupied,
		"freeSlots":    freeSlots,
	})
}

// @Summary      Book a Session
// @Description  Books a session with a doctor (Client only access)
// @Tags         ContentActions
// @Accept       json
// @Produce      json
// @Param        input  body      model.BookingInput  true  "Doctor ID, date, and time"
// @Success      200    {object}  map[string]string  "Booking success message"
// @Failure      400    {object}  map[string]string  "Wrong user or bad input"
// @Failure      500    {object}  map[string]string  "Couldn't book session"
// @Security     ApiKeyAuth
// @Router       /content/client/book [post]
func (h *Handler) bookTimeSlot(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)

	if !isClient {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong user status! (Client/Doctor)",
		})
		return
	}

	var session model.BookingInput
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	booking := model.BookingModel{
		ClientID: userID,
		DoctorID: session.DoctorID,
		Date:     session.Date,
		Time:     session.Time,
		Status:   "booked",
	}

	err := h.service.BookSession(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't book session",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success!",
	})
}
