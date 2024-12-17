package cout

import (
	"encoding/json"
	"fmt"
)

func unmarshalObject(data []byte, result interface{}) error {
	switch result.(type) {
	case *map[string]interface{}:
		return json.Unmarshal(data, result)
	case *interface{}:
		return json.Unmarshal(data, result)
	}
	return json.Unmarshal(data, result)
}

func unmarshalArray(data []byte, result interface{}) error {
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
