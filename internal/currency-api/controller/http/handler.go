package http

import (
	"github.com/agadilkhan/currency-rate/internal/currency-api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.UseCase
}

func NewHandler(service service.UseCase) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(ctx *gin.Context) {

}

func (h *Handler) GetByCode(ctx *gin.Context) {

}

func (h *Handler) ForceUpdate(ctx *gin.Context) {

}
