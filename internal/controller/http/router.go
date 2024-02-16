package http

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		currency := api.Group("/currency")
		{
			currency.GET("/", h.List)
			currency.GET("/:code", h.GetByCode)
			currency.PUT("/force-update", h.ForceUpdate)
		}
	}

	return router
}
