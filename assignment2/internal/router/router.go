package router

import (
	"rest_api_order/internal/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func New() *Router {
	router := Router{router: gin.Default()}
	r := router.router
	r.GET("/orders", controllers.GetAllData)
	r.GET("/orders/:id", controllers.GetSingleData)
	r.POST("/orders/", controllers.InsertData)
	r.PATCH("/orders/:id", controllers.UpdatePATCHMethod)
	r.PUT("/orders/:id", controllers.UpdatePUTMethod)
	r.DELETE("/orders/:id", controllers.DeleteData)

	return &router
}

func (r *Router) StartServer() {
	r.router.Run()
}
