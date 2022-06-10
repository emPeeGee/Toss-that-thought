package handler

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllThoughtsResponse struct {
	Data []entity.Thought `json:"data"`
}

func (h *Handler) test(c *gin.Context) {
	thoughts, err := h.services.Thought.Test()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getAllThoughtsResponse{Data: thoughts.([]entity.Thought)})
}

func (h *Handler) Create(c *gin.Context) {
	var input entity.ThoughtInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Incorrect input",
			"error":   err.Error(),
		})
		return
	}

	if err := h.validate.Struct(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation",
			"error":   err.Error(),
		})
		return
	}

	thoughtMetadata, err := h.services.Thought.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Incorrect server",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}

func (h *Handler) Metadata(c *gin.Context) {
	metadataKey := c.Param("id")

	thoughtMetadata, err := h.services.Thought.Metadata(metadataKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "SQL error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}
