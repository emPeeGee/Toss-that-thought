package handler

import (
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var input entity.ThoughtCreateInput

	if err := c.BindJSON(&input); err != nil {
		flaw.BadRequest(c, "your request looks incorrect", err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		flaw.BadRequest(c, "your request did not pass validation", err.Error())
		return
	}

	thoughtMetadata, err := h.services.Thought.Create(input)
	if err != nil {
		flaw.InternalServer(c, "something went wrong, we are working", err.Error())
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}

func (h *Handler) RetrieveMetadata(c *gin.Context) {
	metadataKey := c.Param("id")

	exists, err := h.services.Thought.CheckMetadataExists(metadataKey)
	if err != nil || !exists {
		flaw.NotFound(c, "such thought does not exists", err.Error())
		return
	}

	thoughtMetadata, err := h.services.Thought.RetrieveMetadata(metadataKey)
	if err != nil {
		flaw.InternalServer(c, "an error encountered during database", err.Error())
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}

func (h *Handler) ThoughtValidity(c *gin.Context) {
	thoughtKey := c.Param("id")

	isValid, err := h.services.Thought.IsThoughtValid(thoughtKey)
	if err != nil || !isValid {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": isValid,
	})
}

func (h *Handler) RetrieveThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	isValid, err := h.services.Thought.IsThoughtValid(thoughtKey)
	if err != nil || !isValid {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	accessThoughtResponse, err := h.services.Thought.RetrieveThought(thoughtKey, accessThoughtInput.Passphrase)
	if err != nil {
		flaw.BadRequest(c, "incorrect password", err.Error())
		return
	}

	c.JSON(http.StatusOK, accessThoughtResponse)
}

// Dublicated code with checking thought if exists
func (h *Handler) BurnThought(c *gin.Context) {
	metadataKey := c.Param("id")

	exists, err := h.services.Thought.CheckMetadataExists(metadataKey)
	if err != nil || !exists {
		flaw.NotFound(c, "such thought does not exists", err.Error())
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	ok, err := h.services.Thought.BurnThought(metadataKey, accessThoughtInput.Passphrase)
	if err != nil || !ok {
		flaw.BadRequest(c, "passphrase is incorrect", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
