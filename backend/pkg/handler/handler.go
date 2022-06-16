package handler

import (
	"github.com/emPeeee/ttt/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.POST("/create", h.Create) // Good to rename endpoint
		api.GET("/metadata/:id", h.RetrieveMetadata)
		api.GET("/thought/:id", h.ThoughtExists)
		api.POST("/thought/:id", h.RetrieveThought)
		api.POST("/thought/:id/burn", h.BurnThought)
	}

	return router
}
