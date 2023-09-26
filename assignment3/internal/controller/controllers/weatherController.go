package controllers

import (
	"assignment_3/internal/controller/validation"
	"assignment_3/internal/repository/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	weather, err := models.GetData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, weather)
}

func UpdateData(c *gin.Context) {
	jsonDecoder := json.NewDecoder(c.Request.Body)
	jsonMap := make(map[string]interface{})
	if err := jsonDecoder.Decode(&jsonMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := validation.ValidateJSONParams(jsonMap, []string{"wind", "water"}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var water float64 = (jsonMap["wind"].(float64))
	var wind float64 = jsonMap["water"].(float64)
	_, err := models.UpdateData(wind, water)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
