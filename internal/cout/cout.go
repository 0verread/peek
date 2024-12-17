package cout

import (
	"fmt"
)

func PrettyPrint(resp []byte) {
	var result interface{}
	err := UnmarshalResp(resp, result)
	if err != nil {
		fmt.Println("Some errror", err)
	}
	fmt.Println(result)
}
