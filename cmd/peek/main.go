package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/0verread/peek/internal/client"
	"github.com/0verread/peek/internal/cout"
	"github.com/spf13/cobra"
)

func buildUrl(u *url.URL) string {
	urlStr := u.String()
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		if strings.HasPrefix(urlStr, "localhost") {
			urlStr = "http://" + urlStr
		} else {
			urlStr = "https://" + urlStr
		}
	}
	return urlStr
}

var rootCmd = &cobra.Command{
	Use:     "peek",
	Version: "0.1.0",
	Short:   "a colorful curl alternative",
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := url.Parse(args[0])

		if err != nil {
			panic(err)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		u, err := url.Parse(args[0])
		if err != nil {
			log.Println("error parsing url", err)
			return err
		}
		urlStr := buildUrl(u)
		verb, _ := cmd.Flags().GetString("verb")
		payload, _ := cmd.Flags().GetString("data")
		header, _ := cmd.Flags().GetString("header")
		response, err := client.Do(urlStr, verb, payload, header)
		if err != nil {
			log.Println("error making request", err)
			return err
		}
		cout.Header(urlStr, verb)
		cout.Stats(response.Status, int(response.Latency))
		cout.PrettyPrint([]byte(response.Body))
		return nil
	},
}

func main() {
	rootCmd.Flags().StringP("verb", "v", "GET", "HTTP verb")
	rootCmd.Flags().StringP("data", "d", "", "data for POST request")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
