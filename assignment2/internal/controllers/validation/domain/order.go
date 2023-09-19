package domain

import (
	"encoding/json"
	"errors"
	"rest_api_order/internal/controllers/validation/general"
	"rest_api_order/internal/repository/models/order"
)

type OrderValidator struct {
}

func (v *OrderValidator) isJSONComplete(jsonByte []byte) error {
	var requiredLen int = 3
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	if len(validationMap) != requiredLen {
		err = errors.New("incomplete JSON: JSON requires customer_name, items, and ordered_at")
	}
	return err
}

func (v *OrderValidator) ContentValidation(jsonByte []byte) (*order.Order, error) {
	var validationStruct order.Order
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	for key := range validationMap {
		if !general.IsInArray(key, []string{"customer_name", "ordered_at", "items"}) {
			err = errors.New("Invalid JSON")
			return &validationStruct, err
		}
	}
	err = json.Unmarshal(jsonByte, &validationStruct)
	return &validationStruct, err
}

// func (v *OrderValidator) DoDuplicateItemsExistInJSON(order *order.Order) error {
// 	var err error
// 	if order.Items == nil {
// 		return err
// 	}
// 	var duplicationExistence bool = false
// 	var itemCodes []string = make([]string, len(order.Items))
// 	for _, item := range order.Items {
// 		if !isInArray(item.Code, itemCodes) {
// 			itemCodes = append(itemCodes, item.Code)
// 		} else {
// 			duplicationExistence = true
// 			break
// 		}
// 	}
// 	if duplicationExistence {
// 		err = errors.New("Duplicate item codes detected")
// 	}
// 	return err
// }

// func (v *OrderValidator) DoesDuplicateExistInDB(order *order.Order) error {
// 	var err error
// 	if item.DoDuplicatesExist(order.Items) {
// 		err = errors.New("code id is a duplicate")
// 	}
// 	return err
// }

func (v *OrderValidator) ValidateJSONStrict(jsonByte []byte) (*order.Order, error) {
	var err error
	var newOrder *order.Order
	err = v.isJSONComplete(jsonByte)
	if err != nil {
		return newOrder, err
	}
	newOrder, err = v.ContentValidation(jsonByte)
	if err != nil {
		return newOrder, err
	}
	return newOrder, err
}
