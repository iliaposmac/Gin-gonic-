package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iliaposmac/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHanlder(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllLists)
			list.GET("/:list_id", h.getListById)
			list.PUT("/:list_id", h.updateList)
			list.DELETE("/:list_id", h.deleteList)
		}

		items := list.Group(":list_id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
