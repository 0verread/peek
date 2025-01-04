package main

import (
	"errors"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"

	"github.com/0verread/peek/internal/client"
	"github.com/0verread/peek/internal/cout"
)

var rootCmd = &cobra.Command{
	Use:     "peek",
	Version: "0.1.0",
	Short:   "test apis better way",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires an argument")
		}

		url, err := url.Parse(args[0])

		if err != nil {
			panic(err)
		}

		port := url.Port()

		if url.Scheme == "http" && port == "" {
			log.Println("Need port for http request")
			os.Exit(0)
		}

		if url.Scheme == "https" && port == "" {
			port = "441"
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		verb, _ := cmd.Flags().GetString("verb")
		payload, _ := cmd.Flags().GetString("data")
		header, _ := cmd.Flags().GetString("header")
		response, _ := client.Do(args[0], verb, payload, header)
		cout.Header(args[0], verb)
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
