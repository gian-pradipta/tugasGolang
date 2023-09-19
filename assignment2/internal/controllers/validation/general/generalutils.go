package general

import (
	"encoding/json"
	"errors"
	"strconv"
)

func ValidateParam(id string) (uint, error) {
	var intId int
	var err error
	intId, err = strconv.Atoi(id)
	return uint(intId), err
}

func IsJSONComplete(jsonByte []byte) error {
	var requiredLen int = 3
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	if len(validationMap) != requiredLen {
		err = errors.New("incomplete JSON: JSON requires customer_name, items, and ordered_at")
	}
	return err
}
func IsInArray(word string, words []string) bool {
	var result bool = false
	for i := range words {
		if words[i] == word {
			result = true
			break
		}
	}
	return result
}
