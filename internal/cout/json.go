package cout

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func unmarshalObject(data []byte, result interface{}) error {
	switch result.(type) {
	case map[string]interface{}:
		fmt.Println(json.Unmarshal(data, result))
		return json.Unmarshal(data, result)
	case interface{}:
		return json.Unmarshal(data, result)
	}
	return json.Unmarshal(data, result)
}

func unmarshalArray(data []byte, result interface{}) error {
	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr {
		return fmt.Errorf("result must be a pointer")
	}

	// Get the underlying type
	elemType := resultType.Elem()

	// Check if it's a slice
	if elemType.Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}

	// Create a new slice of the correct type
	sliceType := reflect.SliceOf(elemType.Elem())
	slice := reflect.New(sliceType).Elem()

	// Unmarshal into the new slice
	if err := json.Unmarshal(data, slice.Addr().Interface()); err != nil {
		return fmt.Errorf("failed to unmarshal array: %v", err)
	}

	// Set the original result to the new slice
	reflect.ValueOf(result).Elem().Set(slice)
	return nil
}

// func to unmarshal the response based on response type
// response could be either a single object or array of objects
func UnmarshalResp(resp []byte, result interface{}) error {
	var rawData interface{}
	if err := json.Unmarshal(resp, &rawData); err != nil {
		return fmt.Errorf("failed to parse Json: %v", err)
	}
	switch rawData.(type) {
	case []interface{}:
		return unmarshalArray(resp, result)
	case map[string]interface{}:
		return unmarshalObject(resp, result)
	default:
		// do nothing and return resp when invalid data type
		return nil
	}
}
