package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nvsces/flw-server-go/pkg/service"
)

type Handler struct{
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services:services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
	}

	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.POST("/create", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}