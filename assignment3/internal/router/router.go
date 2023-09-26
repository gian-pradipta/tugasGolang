package router

import (
	"assignment_3/internal/controller/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.New()

func routing() {
	router.GET("/weather", controllers.GetData)
	router.PUT("/weather", controllers.UpdateData)
}

func Run() error {
	routing()
	err := router.Run("localhost:8000")
	return err
}
