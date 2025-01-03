package cout

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/0verread/peek/pkg/prettify"
)

func Stats(status int, latency int) {
	coloredStatus := prettify.Status(status)
	coloredLatency := prettify.Latency(latency)
	coloredStats := fmt.Sprintf("Status: %s  Time Taken: %s ms\n", coloredStatus(status), coloredLatency(latency))
	os.Stdout.Write([]byte(coloredStats))
}

func Header(url string) {
	coloredUrl, _ := prettify.Prettify(url)
	// coloredFormattedUrl := fmt.Sprintf("Url: %v\n", coloredUrl)
	os.Stdout.Write(coloredUrl)
}

func Status(status int) {
	coloredStatus := prettify.Status(status)
	os.Stdout.Write([]byte(coloredStatus(status)))
}

func Latency(latency int) {
	coloredLatency := prettify.Latency(latency)
	os.Stdout.Write([]byte(coloredLatency(latency)))
}

func PrettyPrint(resp []byte) {
	var err error
	var coloredRespBody []byte
	var rawData interface{}
	if err = json.Unmarshal(resp, &rawData); err != nil {
		log.Println("failed to parse Json: ", err)
	}

	switch rawData.(type) {
	case []interface{}:
		var result []map[string]interface{}
		err = UnmarshalResp(resp, &result)

		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}
		coloredRespBody, err = prettify.Prettify(result)
	case map[string]interface{}:
		var result map[string]interface{}
		err = UnmarshalResp(resp, &result)
		fmt.Println("result: ", result)
		if err != nil {
			log.Println("Error in Unmarshal Response, error: ", err)
		}
		coloredRespBody, err = prettify.Prettify(result)
		fmt.Println("coloredRespBody: ", coloredRespBody)
	}
	if err != nil {
		log.Println("Error in colored Response", err)
	}
	os.Stdout.Write(coloredRespBody)
}
