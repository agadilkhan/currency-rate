package http

import (
	"github.com/agadilkhan/currency-rate/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	res, err := h.service.List(ctx)
	if err != nil {
		log.Printf("failed to List err: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	res, err := h.service.GetByCode(ctx, code)
	if err != nil {
		log.Printf("failed to GetByCode err: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) ForceUpdate(ctx *gin.Context) {
	err := h.service.Update(ctx)
	if err != nil {
		log.Printf("failed to ForceUpdate err: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
