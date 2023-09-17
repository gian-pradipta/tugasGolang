package controllers

import (
	"io"
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
	err := c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	latestId := order.InsertData(&newOrder)
	c.JSON(http.StatusOK, gin.H{
		"order": order.GetSingleData(latestId),
	})

}

func DeleteData(c *gin.Context) {
	var idString string = c.Param("id")
	var id int
	var err error
	id, err = strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	deletedOrder := order.GetSingleData(uint(id))
	err = order.DeleteData(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": deletedOrder,
	})
}
func UpdatePATCHMethod(c *gin.Context) {
	var notValidatedId string = c.Param("id")
	var validatedParam uint
	var err error
	if validatedParam, err = validateParam(notValidatedId); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var jsonByte []byte
	if jsonByte, err = io.ReadAll(c.Request.Body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var newOrder *order.Order
	if newOrder, err = contentValidation(jsonByte); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = order.UpdateOrder(validatedParam, newOrder); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	GetSingleData(c)
}
func UpdatePUTMethod(c *gin.Context) {
	var notValidatedId string = c.Param("id")
	var err error
	var validatedParam uint
	if validatedParam, err = validateParam(notValidatedId); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var jsonByte []byte
	if jsonByte, err = io.ReadAll(c.Request.Body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var newOrder *order.Order
	if newOrder, err = validateLengthAndContent(jsonByte); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = order.UpdateOrder(validatedParam, newOrder); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	GetSingleData(c)
}
