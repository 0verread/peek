package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/0verread/peek/internal/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "peek",
	Version: "0.1.0",
	Short:   "test apis better way",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires an argument")
		}

		fmt.Println("url: ", args[0])
		url, err := url.Parse(args[0])

		if err != nil {
			panic(err)
		}

		port := url.Port()

		if url.Scheme == "http" && port == "" {
			fmt.Println("Need port for http request")
			os.Exit(0)
		}

		if url.Scheme == "https" && port == "" {
			port = "441"
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running url: ", args[0])
		verb, _ := cmd.Flags().GetString("verb")

		payload, _ := cmd.Flags().GetString("data")

		header, _ := cmd.Flags().GetString("header")

		client.Do(args[0], verb, payload, header)
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
