package controllers

import (
	"net/http"
	"rest_api_order/internal/repository/models/order"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"orders": order.GetAllData(),
	})
}

func GetSingleData(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var err error
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order": order.GetSingleData(uint(intId)),
	})
}

func InsertData(c *gin.Context) {
	var newOrder order.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.InsertData(&newOrder)

}
