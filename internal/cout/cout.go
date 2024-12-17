package cout

import (
	"fmt"

	"github.com/0verread/peek/pkg/prettyjson"
)

func PrettyPrintString(str string) {
	coloredStr, _ := prettyjson.Prettify(str)
	fmt.Println(string(coloredStr))
}

func PrettyPrint(resp []byte) {
	var result []map[string]interface{}
	err := UnmarshalResp(resp, &result)
	if err != nil {
		fmt.Println("Some errror", err)
	}
	// fmt.Println(result)
	coloredRespBody, _ := prettyjson.Prettify(result)
	fmt.Println(string(coloredRespBody))
}
