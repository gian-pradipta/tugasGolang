package controllers

import (
	"encoding/json"
	"errors"
	"rest_api_order/internal/repository/models/order"
	"strconv"
)

func validateJSON(m map[string]interface{}) bool {
	qualified := true
	if len(m) != 3 {
		qualified = false
	}
	return qualified
}

func validateParam(id string) (uint, error) {
	var intId int
	var err error
	intId, err = strconv.Atoi(id)
	return uint(intId), err
}

func lengthValidation(jsonByte []byte) error {
	var requiredLen int = 3
	var validationMap map[string]interface{}
	var err error
	json.Unmarshal(jsonByte, &validationMap)
	if len(validationMap) != requiredLen {
		err = errors.New("Invalid JSON request")
	}
	return err
}
func contentValidation(jsonByte []byte) (*order.Order, error) {
	var validationStruct order.Order
	var err error
	err = json.Unmarshal(jsonByte, &validationStruct)
	return &validationStruct, err
}
func validateLengthAndContent(jsonByte []byte) (*order.Order, error) {
	var err error
	var newOrder *order.Order
	err = lengthValidation(jsonByte)
	if err != nil {
		return newOrder, err
	}
	newOrder, err = contentValidation(jsonByte)
	if err != nil {
		return newOrder, err
	}
	return newOrder, err
}
