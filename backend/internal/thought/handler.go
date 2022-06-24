package thought

import (
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, validate *validator.Validate, logger log.Logger) {
	h := handler{service, logger, validate}

	api := r.Group("")
	{
		api.POST("/create", h.create) // Good to rename endpoint
		api.GET("/metadata/:id", h.retrieveMetadata)
		api.GET("/thought/:id", h.thoughtValidity)
		api.POST("/thought/:id", h.retrieveThought)
		api.POST("/thought/:id/burn", h.burnThought)
	}
}

type handler struct {
	service  Service
	logger   log.Logger
	validate *validator.Validate
}

func (h *handler) create(c *gin.Context) {
	var input entity.ThoughtCreateInput

	if err := c.BindJSON(&input); err != nil {
		flaw.BadRequest(c, "your request looks incorrect", err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		flaw.BadRequest(c, "your request did not pass validation", err.Error())
		return
	}

	thoughtMetadata, err := h.service.Create(input)
	if err != nil {
		flaw.InternalServer(c, "something went wrong, we are working", err.Error())
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}

func (h *handler) retrieveMetadata(c *gin.Context) {
	metadataKey := c.Param("id")

	exists, err := h.service.CheckMetadataExists(metadataKey)
	if err != nil || !exists {
		flaw.NotFound(c, "such thought does not exists", err.Error())
		return
	}

	thoughtMetadata, err := h.service.RetrieveMetadata(metadataKey)
	if err != nil {
		flaw.InternalServer(c, "an error encountered during database", err.Error())
		return
	}

	c.JSON(http.StatusOK, thoughtMetadata)
}

func (h *handler) thoughtValidity(c *gin.Context) {
	thoughtKey := c.Param("id")

	isValid, err := h.service.IsThoughtValid(thoughtKey)
	if err != nil || !isValid {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": isValid,
	})
}

func (h *handler) retrieveThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	isValid, err := h.service.IsThoughtValid(thoughtKey)
	if err != nil || !isValid {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	accessThoughtResponse, err := h.service.RetrieveThought(thoughtKey, accessThoughtInput.Passphrase)
	if err != nil {
		flaw.BadRequest(c, "incorrect password", err.Error())
		return
	}

	c.JSON(http.StatusOK, accessThoughtResponse)
}

// Dublicated code with checking thought if exists
func (h *handler) burnThought(c *gin.Context) {
	metadataKey := c.Param("id")

	exists, err := h.service.CheckMetadataExists(metadataKey)
	if err != nil || !exists {
		flaw.NotFound(c, "such thought does not exists", err.Error())
		return
	}

	var accessThoughtInput entity.ThoughtPassphraseInput
	if err := c.BindJSON(&accessThoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	ok, err := h.service.BurnThought(metadataKey, accessThoughtInput.Passphrase)
	if err != nil || !ok {
		flaw.BadRequest(c, "passphrase is incorrect", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": ok,
	})
}
