package handler

import (
	"github.com/emPeeee/ttt/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
}

func NewHandler(services *service.Service, validator *validator.Validate) *Handler {
	return &Handler{services: services, validate: validator}
}

func (h *Handler) InitializeRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/create", h.Create)
		api.GET("/metadata/:id", h.Metadata)
		api.GET("/thought/:id", h.ThoughtExists)
		api.POST("/thought/:id", h.AccessThought)
	}

	return router
}
