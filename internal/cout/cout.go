package cout

import (
	"encoding/json"
	"fmt"
	"github.com/0verread/peek/pkg/prettyjson"
	"log"
	"os"
)

func PrettyPrintString(str string) {
	coloredStr, _ := prettyjson.Prettify(str)
	log.Println(string(coloredStr))
}

func PrettyPrint(resp []byte) {
	var err error
	var rawData interface{}
	if err := json.Unmarshal(resp, &rawData); err != nil {
		fmt.Println("failed to parse Json: ", err)
	}

	switch rawData.(type) {
	case []interface{}:
		var result []map[string]interface{}
		// err = UnmarshalResp(resp, &result)
		err = UnmarshalResp(resp, &result)
		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}
		coloredRespBody, err := prettyjson.Prettify(result)
		if err != nil {
			log.Println("Error in colored Response", err)
		}
		os.Stdout.Write(coloredRespBody)
	case map[string]interface{}:
		var result map[string]interface{}
		// err = UnmarshalResp(resp, &result)
		err = UnmarshalResp(resp, &result)
		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}

		coloredRespBody, err := prettyjson.Prettify(result)
		if err != nil {
			log.Println("Error in colored Response", err)
		}
		os.Stdout.Write(coloredRespBody)
	}

}
