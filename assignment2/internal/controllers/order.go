package controllers

import (
	"io"
	"net/http"
	"rest_api_order/internal/repository/models/order"

	"github.com/gin-gonic/gin"
)

func errToJSON(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func ShowAllData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"orders": order.GetAllData(),
	})
}

func ShowSingleData(ctx *gin.Context) {
	var notValidatedId string = ctx.Param("id")
	var err error
	validatedId, err := validateParam(notValidatedId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	jsonData, err := order.GetSingleData(validatedId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errToJSON(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order": jsonData,
	})
}

func CreateData(c *gin.Context) {
	var jsonByte []byte
	var err error
	jsonByte, err = io.ReadAll(c.Request.Body)

	var newOrder *order.Order
	if newOrder, err = validateJSONStrict(jsonByte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	if err = doDuplicateItemsExistInJSON(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	if err = DoesDuplicateExistInDB(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	latestId := order.InsertData(newOrder)
	jsonData, _ := order.GetSingleData(latestId)
	c.JSON(http.StatusOK, gin.H{
		"order": jsonData,
	})

}

func DeleteData(c *gin.Context) {
	var notValidatedId string = c.Param("id")
	var id uint
	var err error

	if id, err = validateParam(notValidatedId); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errToJSON(err))
		return
	}
	deletedOrder, err := order.DeleteData(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, errToJSON(err))
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
		c.AbortWithStatusJSON(http.StatusNotFound, errToJSON(err))
		return
	}
	var jsonByte []byte
	if jsonByte, err = io.ReadAll(c.Request.Body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	var newOrder *order.Order
	if newOrder, err = contentValidation(jsonByte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	if err = doDuplicateItemsExistInJSON(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	if err = DoesDuplicateExistInDB(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	var updatedOrder *order.Order
	if updatedOrder, err = order.UpdateOrder(validatedParam, newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": updatedOrder,
	})

}
func UpdatePUTMethod(c *gin.Context) {
	var notValidatedId string = c.Param("id")
	var err error

	var validatedParam uint
	if validatedParam, err = validateParam(notValidatedId); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errToJSON(err))
		return
	}

	var jsonByte []byte
	if jsonByte, err = io.ReadAll(c.Request.Body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	var newOrder *order.Order
	if newOrder, err = validateJSONStrict(jsonByte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	if err = doDuplicateItemsExistInJSON(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	if err = DoesDuplicateExistInDB(newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}

	var updatedOrder *order.Order
	if updatedOrder, err = order.UpdateOrder(validatedParam, newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errToJSON(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": updatedOrder,
	})
}
