package auth

import (
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
		api.POST("/signUp", h.signUp)
		api.POST("/signIn", h.signIn)
	}
}

type handler struct {
	service  Service
	logger   log.Logger
	validate *validator.Validate
}

func (h *handler) signUp(c *gin.Context) {
	var input createUserDTO

	if err := c.BindJSON(&input); err != nil {
		flaw.BadRequest(c, "your request looks incorrect", err.Error())
		return
	}

	h.logger.Debug(input)

	if err := h.validate.Struct(input); err != nil {
		flaw.BadRequest(c, "your request did not pass validation", err.Error())
		return
	}

	err := h.service.createUser(input)
	if err != nil {
		flaw.InternalServer(c, "something went wrong, we are working", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}

func (h *handler) signIn(c *gin.Context) {
	var credentials credentialsDTO

	if err := c.BindJSON(&credentials); err != nil {
		flaw.BadRequest(c, "incorrect body", err.Error())
		return
	}

	if err := h.validate.Struct(credentials); err != nil {
		flaw.BadRequest(c, "incorrect body", err.Error())
		return
	}

	token, err := h.service.generateToken(credentials)
	if err != nil {
		flaw.InternalServer(c, "something went wrong, we are working", err.Error())
		return
	}

	h.logger.Debug(token)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
