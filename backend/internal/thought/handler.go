package thought

import (
	"github.com/emPeeee/ttt/internal/auth"
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func RegisterHandlers(r *gin.RouterGroup, service Service, validate *validator.Validate, logger log.Logger) {
	h := handler{service, logger, validate}

	api := r.Group("")
	{
		api.POST("/create", h.create)
		api.GET("/metadata/:id", h.retrieveMetadata)
		api.GET("/thought/:id", h.retrieveThoughtPassphraseInfo)
		api.POST("/thought/:id", h.retrieveThought)
		api.POST("/thought/:id/burn", h.burnThought)
	}
}

type handler struct {
	service  Service
	logger   log.Logger
	validate *validator.Validate
}

// such user already exists?
func (h *handler) create(c *gin.Context) {
	var input CreateDTO

	// TODO: handle authenticated users
	userId, err := auth.GetUserId(c)
	if err != nil {
		h.logger.Debug(err.Error(), "Not authenticated")
	} else {
		h.logger.Debug(userId, "Authenticated")
	}

	if err := c.BindJSON(&input); err != nil {
		flaw.BadRequest(c, "your request looks incorrect", err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		flaw.BadRequest(c, "your request did not pass validation", err.Error())
		return
	}

	// TODO: Future, to make a validator for this https://github.com/go-playground/validator/issues/494
	maxLifetime := time.Now().AddDate(0, 0, 7)
	if input.Lifetime.After(maxLifetime) {
		flaw.BadRequest(c, "lifetime cannot be so long", "lifetime cannot be more that 7 days")
		return
	}

	createdThought, err := h.service.Create(input)
	if err != nil {
		flaw.InternalServer(c, "something went wrong, we are working", err.Error())
		return
	}

	c.JSON(http.StatusOK, createdThought)
}

func (h *handler) retrieveMetadata(c *gin.Context) {
	metadataKey := c.Param("id")

	err := h.service.CheckMetadataExists(metadataKey)
	if err != nil {
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

func (h *handler) retrieveThoughtPassphraseInfo(c *gin.Context) {
	thoughtKey := c.Param("id")

	info, err := h.service.RetrieveThoughtPassphraseInfo(thoughtKey)
	if err != nil {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	c.JSON(http.StatusOK, info)
}

func (h *handler) retrieveThought(c *gin.Context) {
	thoughtKey := c.Param("id")

	err := h.service.IsThoughtValid(thoughtKey)
	if err != nil {
		flaw.NotFound(c, "thought it either never existed or already has been viewed", err.Error())
		return
	}

	var thoughtInput PassphraseDTO
	if err := c.BindJSON(&thoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	thoughtResponse, err := h.service.RetrieveThoughtByPassphrase(thoughtKey, thoughtInput.Passphrase)
	if err != nil {
		flaw.BadRequest(c, "incorrect password", err.Error())
		return
	}

	c.JSON(http.StatusOK, thoughtResponse)
}

// Dublicated code with checking thought if exists
func (h *handler) burnThought(c *gin.Context) {
	metadataKey := c.Param("id")

	err := h.service.CheckMetadataExists(metadataKey)
	if err != nil {
		flaw.NotFound(c, "such thought does not exists", err.Error())
		return
	}

	var thoughtInput PassphraseDTO
	if err := c.BindJSON(&thoughtInput); err != nil {
		flaw.BadRequest(c, "your request seems to be incorrect", err.Error())
		return
	}

	err = h.service.BurnThought(metadataKey, thoughtInput.Passphrase)
	if err != nil {
		flaw.BadRequest(c, "passphrase is incorrect", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
