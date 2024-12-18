package cout

import (
	"encoding/json"
	"fmt"
)

type unmarshalArrayResp []map[string]interface{}

func unmarshalObject(data []byte, result interface{}) error {
	switch result.(type) {
	case map[string]interface{}:
		return json.Unmarshal(data, result)
	case interface{}:
		return json.Unmarshal(data, result)
	}
	return json.Unmarshal(data, result)
}

func unmarshalArray(data []byte, result interface{}) error {
	// Ensure result is a pointer to a slice of maps
	slicePtr, ok := result.(*[]map[string]interface{})
	if !ok {
		return fmt.Errorf("result must be a pointer to []map[string]interface{}")
	}

	// Unmarshal directly into the slice of maps
	return json.Unmarshal(data, slicePtr)
}

// func to unmarshal the response based on response type
// response could be either a single object or array of objects
func UnmarshalResp(resp []byte, result interface{}) error {
	var rawData interface{}
	if err := json.Unmarshal(resp, &rawData); err != nil {
		fmt.Println("failed to parse Json: ", err)
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
