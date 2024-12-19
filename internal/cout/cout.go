package cout

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/0verread/peek/pkg/prettyjson"
)

func PrettyPrintString(str string) {
	coloredStr, _ := prettyjson.Prettify(str)
	os.Stdout.Write(coloredStr)
}

func PrettyPrintStatus(status int) {
	coloredStatus, _ := prettyjson.Prettify(status)
	os.Stdout.Write(coloredStatus)
}

func PrettyPrint(resp []byte) {
	var err error
	var coloredRespBody []byte
	var rawData interface{}
	if err = json.Unmarshal(resp, &rawData); err != nil {
		fmt.Println("failed to parse Json: ", err)
	}

	switch rawData.(type) {
	case []interface{}:
		var result []map[string]interface{}
		err = UnmarshalResp(resp, &result)
		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}
		coloredRespBody, err = prettyjson.Prettify(result)
	case map[string]interface{}:
		var result map[string]interface{}
		err = UnmarshalResp(resp, &result)
		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}
		coloredRespBody, err = prettyjson.Prettify(result)
	}
	if err != nil {
		log.Println("Error in colored Response", err)
	}
	os.Stdout.Write(coloredRespBody)
}
