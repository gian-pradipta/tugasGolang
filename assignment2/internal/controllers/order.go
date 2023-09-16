package controllers

import (
	"encoding/json"
	"errors"
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
		return
	}

	jsonByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var validationMap map[string]interface{}
	var newOrder order.Order
	json.Unmarshal(jsonByte, &validationMap)
	err = json.Unmarshal(jsonByte, &newOrder)
	if !validateJSON(validationMap) {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad Request"))
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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

func UpdatePartOfOrder(c *gin.Context) {
	var idString string = c.Param("id")
	var id int
	var err error
	id, err = strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var newOrder order.Order
	err = c.ShouldBindJSON(&newOrder)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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

func validateJSON(m map[string]interface{}) bool {
	qualified := true
	if len(m) != 3 {
		qualified = false
	}
	return qualified
}
