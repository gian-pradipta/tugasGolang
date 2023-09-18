package controllers

import (
	"encoding/json"
	"errors"
	"rest_api_order/internal/repository/models/order"
	"strconv"
)

func validateParam(id string) (uint, error) {
	var intId int
	var err error
	intId, err = strconv.Atoi(id)
	return uint(intId), err
}

func isJSONComplete(jsonByte []byte) error {
	var requiredLen int = 3
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	if len(validationMap) != requiredLen {
		err = errors.New("Invalid JSON request")
	}
	return err
}
func isInArray(word string, words []string) bool {
	var result bool = false
	for i := range words {
		if words[i] == word {
			result = true
			break
		}
	}
	return result
}
func contentValidation(jsonByte []byte) (*order.Order, error) {
	var validationStruct order.Order
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	for key := range validationMap {
		if !isInArray(key, []string{"customer_name", "ordered_at", "items"}) {
			err = errors.New("Invalid JSON")
			return &validationStruct, err
		}
	}
	err = json.Unmarshal(jsonByte, &validationStruct)
	return &validationStruct, err
}

func validateDuplicateItems(order *order.Order) error {
	var err error
	if order.Items == nil {
		return err
	}
	var duplicationExistence bool = false
	var itemCodes []string = make([]string, len(order.Items))
	for _, item := range order.Items {
		if !isInArray(item.Code, itemCodes) {
			itemCodes = append(itemCodes, item.Code)
		} else {
			duplicationExistence = true
			break
		}
	}
	if duplicationExistence {
		err = errors.New("Duplicate item codes detected")
	}
	return err
}

func validateJSONStrict(jsonByte []byte) (*order.Order, error) {
	var err error
	var newOrder *order.Order
	err = isJSONComplete(jsonByte)
	if err != nil {
		return newOrder, err
	}
	newOrder, err = contentValidation(jsonByte)
	if err != nil {
		return newOrder, err
	}
	return newOrder, err
}
