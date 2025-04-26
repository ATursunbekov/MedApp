package handler

import (
	"MedApp/internal/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// createMedicine godoc
// @Summary Create a new medicine
// @Description Creates a new medicine entry in the database
// @Tags Medicine
// @Accept json
// @Produce json
// @Param medicine body model.Medicine true "Medicine data"
// @Success 201 {object} map[string]interface{} "message: Medicine created successfully, id: medicine.ID"
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 500 {object} map[string]string "error: Failed to create medicine"
// @Router /medicine [post]
func (h *Handler) createMedicine(c *gin.Context) {
	var medicine model.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := h.service.Medicine.Create(c.Request.Context(), &medicine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create medicine: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Medicine created successfully",
		"data":    gin.H{"id": medicine.ID},
	})
}

// getMedicineByID godoc
// @Summary Get medicine by ID
// @Description Retrieves a medicine entry by its ID
// @Tags Medicine
// @Produce json
// @Param id path string true "Medicine ID"
// @Success 200 {object} model.Medicine
// @Failure 400 {object} map[string]string "error: Invalid ID"
// @Failure 404 {object} map[string]string "error: Medicine not found"
// @Failure 500 {object} map[string]string "error: Failed to get medicine"
// @Router /medicine/{id} [get]
func (h *Handler) getMedicineByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	medicine, err := h.service.Medicine.GetByID(c.Request.Context(), id)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get medicine: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}

// getAllMedicines godoc
// @Summary Get all medicines
// @Description Retrieves a list of all medicines
// @Tags Medicine
// @Produce json
// @Success 200 {array} model.Medicine
// @Failure 500 {object} map[string]string "error: Failed to get medicines"
// @Router /medicine [get]
func (h *Handler) getAllMedicines(c *gin.Context) {
	medicines, err := h.service.Medicine.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get medicines: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicines})
}

// deleteMedicine godoc
// @Summary Delete medicine by ID
// @Description Deletes a medicine entry by its ID
// @Tags Medicine
// @Produce json
// @Param id path string true "Medicine ID"
// @Success 204
// @Failure 400 {object} map[string]string "error: Invalid ID"
// @Failure 404 {object} map[string]string "error: Medicine not found"
// @Failure 500 {object} map[string]string "error: Failed to delete medicine"
// @Router /medicine/{id} [delete]
func (h *Handler) deleteMedicine(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Medicine.Delete(c.Request.Context(), id); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medicine: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
