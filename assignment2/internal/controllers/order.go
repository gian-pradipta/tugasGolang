package controllers

import (
	"net/http"
	"rest_api_order/internal/repository/models/order"
	"strconv"
	"time"

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
	err := c.ShouldBindJSON(&newOrder)
	newOrder.OrderedAt = time.Now()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.InsertData(&newOrder)

}

func DeleteData(c *gin.Context) {
	var idString string = c.Param("id")
	var id int
	var err error
	id, err = strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = order.DeleteData(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdateAnEntireOrder(c *gin.Context) {
	var idString string = c.Param("id")
	var id int
	var err error
	id, err = strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var newOrder order.Order
	err = c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = order.UpdateAnEntireOrder(uint(id), &newOrder)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

}

func UpdatePartOfOrder(id uint, order order.Order) {

}
