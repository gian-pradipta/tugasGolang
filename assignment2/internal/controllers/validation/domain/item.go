package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"rest_api_order/internal/controllers/validation/general"
	"rest_api_order/internal/repository/models/item"
	"rest_api_order/internal/repository/models/order"
)

type ItemValidator struct {
}

func (v *ItemValidator) isSingleJSONComplete(item item.Item) bool {

	var result bool = true
	result = result && !(item.Code == "")
	result = result && !(item.Description == "")
	result = result && !(item.Quantity == 0)

	return result
}

func (v *ItemValidator) isJSONComplete(jsonByte []byte) error {
	// var validationMap map[string]interface{}
	var validationStruct order.Order
	var err error
	err = json.Unmarshal(jsonByte, &validationStruct)
	if err != nil {
		return err
	}
	for _, item := range validationStruct.Items {
		if !v.isSingleJSONComplete(item) {
			err = errors.New("incomplete in order's items JSON: items requires code, decription, and quantity")
			return err
		}
	}
	return err
}
func (v *ItemValidator) ContentValidation(jsonByte []byte) (*item.Item, error) {
	var validationStruct item.Item
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)

	for key := range validationMap {
		if !general.IsInArray(key, []string{"code", "description", "quantity"}) {
			fmt.Println(key)
			err = errors.New("Invalid JSON 2")
			return &validationStruct, err
		}
	}
	err = json.Unmarshal(jsonByte, &validationStruct)
	return &validationStruct, err
}

func (v *ItemValidator) DoDuplicateItemsExistInJSON(items []item.Item) error {
	var err error
	if items == nil {
		return err
	}
	var duplicationExistence bool = false
	var itemCodes []string = make([]string, len(items))
	for _, item := range items {
		if !general.IsInArray(item.Code, itemCodes) {
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

func (v *ItemValidator) DoesDuplicateExistInDB(items []item.Item) error {
	var err error
	if item.DoDuplicatesExist(items) {
		err = errors.New("code id is a duplicate")
	}
	return err
}

func (v *ItemValidator) ValidateJSONStrict(jsonByte []byte) (*item.Item, error) {
	var err error
	var newItem *item.Item
	err = v.isJSONComplete(jsonByte)
	if err != nil {
		return newItem, err
	}

	return newItem, err
}
