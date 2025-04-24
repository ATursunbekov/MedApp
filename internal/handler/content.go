package handler

import "github.com/gin-gonic/gin"

// @Summary      Check api workflow
// @Description  Returns saved user ID and client status from context (set by middleware)
// @Security     ApiKeyAuth
// @Tags         ContentCheck
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "user ID and status"
// @Router       /content/check [get]
func (h *Handler) login(c *gin.Context) {
	userID := c.GetString(savedID)
	isClient := c.GetBool(status)

	c.JSON(200, gin.H{
		"ID":     userID,
		"status": isClient,
	})
}

// @Summary      Get All Cat Facts
// @Description  Returns a list of cat facts from the database or service
// @Tags         ContentCheck
// @Produce      json
// @Success      200  {array}   model.CatFact
// @Failure      500  {object}  map[string]string  "Internal error"
// @Router       /getCatFacts [get]
func (h *Handler) getCatFacts(c *gin.Context) {
	facts, err := h.service.GetCatFacts()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, facts)
}

// @Summary      Get Single Cat Fact
// @Description  Returns a cat fact by its ID
// @Tags         ContentCheck
// @Produce      json
// @Param        id   path      string  true  "Cat Fact ID"
// @Success      200  {object}  model.CatFact
// @Failure      400  {object}  map[string]string  "Missing ID"
// @Failure      500  {object}  map[string]string  "Internal error"
// @Router       /catfacts/{id} [get]
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
