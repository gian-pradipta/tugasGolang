package validation

import "errors"

func isInArray(keyword string, words []string) bool {
	var result bool = true
	comparitionMap := make(map[string]int)
	for _, word := range words {
		comparitionMap[word] = 1
	}
	if comparitionMap[keyword] == 0 {
		result = false
	}
	return result
}

func ValidateJSONParams(unvalidatedMap map[string]interface{}, allowedParams []string) error {
	var err error

	if len(unvalidatedMap) != len(allowedParams) {
		err = errors.New("Invalid JSON elements")
		return err
	}

	for key := range unvalidatedMap {
		if !isInArray(key, allowedParams) {
			err = errors.New("Invalid JSON parameter")
			return err
		}
	}

	return err
}
