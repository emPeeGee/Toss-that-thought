package handler

import (
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var input entity.ThoughtCreateInput

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

func (h *Handler) RetrieveMetadata(c *gin.Context) {
	metadataKey := c.Param("id")

	thoughtMetadata, err := h.services.Thought.RetrieveMetadata(metadataKey)
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

func (h *Handler) RetrieveThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	exists, err := h.services.Thought.CheckThoughtExists(thoughtKey)
	if err != nil || !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	accessThoughtResponse, err := h.services.Thought.RetrieveThought(thoughtKey, accessThoughtInput.Passphrase)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Incorrect password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accessThoughtResponse)
}

// Dublicated code with checking thought if exists
func (h *Handler) BurnThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	exists, err := h.services.Thought.CheckThoughtExists(thoughtKey)
	if err != nil || !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Such thought does not exists",
			"error":   err.Error(),
		})
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Incorrect body",
			"error":   err.Error(),
		})
		return
	}

	ok, err := h.services.Thought.BurnThought(thoughtKey, accessThoughtInput.Passphrase)
	if err != nil || !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Passphrase is incorrect",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
