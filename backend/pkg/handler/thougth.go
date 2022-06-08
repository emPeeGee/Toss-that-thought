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
