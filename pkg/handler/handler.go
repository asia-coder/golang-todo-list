package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	var router = gin.New()

	var auth = router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	var api = router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}
		}
	}

	return router
}
