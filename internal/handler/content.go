package handler

import "github.com/gin-gonic/gin"

func (h *Handler) login(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)
	c.JSON(200, gin.H{
		"ID":     userID,
		"status": isClient,
	})
}

func (h *Handler) getCatFacts(c *gin.Context) {
	facts, err := h.service.GetCatFacts()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, facts)
}

func (h *Handler) getCatFact(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Missing id in path"})
		return
	}

	facts, err := h.service.GetCatFact(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	c.JSON(200, facts)
}
