package handler

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// ThoughtExists I don't really like this name
// But i frontend I should say that thought does not exist or is burned
func (h *Handler) ThoughtExists(c *gin.Context) {
	thoughtKey := c.Param("id")

	exists, err := h.services.Thought.CheckThoughtExists(thoughtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": exists,
	})
}

func (h *Handler) AccessThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	exists, err := h.services.Thought.CheckThoughtExists(thoughtKey)
	if err != nil || !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	var accessThoughtInput entity.AccessThoughtInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	accessThoughtResponse, err := h.services.Thought.AccessThought(thoughtKey, accessThoughtInput.Passphrase)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Incorrect password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accessThoughtResponse)
}