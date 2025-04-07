package handler

import (
	"log/slog"
	"net/http"

	"github.com/fentezi/unit-converter/model"
	"github.com/fentezi/unit-converter/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	log     *slog.Logger
	service *service.Service
}

func NewHandler(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) Convert(c *gin.Context) {
	var request model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.Error("request error", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "fill in all fields",
		})
	}

	response := h.service.Convert(request)

	c.JSON(http.StatusOK, response)
}
