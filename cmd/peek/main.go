package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/0verread/peek/internal/client"
	"github.com/0verread/peek/internal/cout"
	"github.com/spf13/cobra"
)

func buildUrl(u *url.URL) *url.URL {
	if u.Scheme == "" || u.Host == "" {
		if strings.HasPrefix(u.Host, "localhost") || strings.HasPrefix(u.Scheme, "localhost") {
			return &url.URL{
				Scheme:      "http",
				Host:        "localhost",
				Path:        u.Path,
				RawQuery:    u.RawQuery,
				RawFragment: u.RawFragment,
			}
		} else {
			return &url.URL{
				Scheme:      "https",
				Host:        u.Host,
				Path:        u.Path,
				RawQuery:    u.RawQuery,
				RawFragment: u.RawFragment,
			}
		}
	}
	return u
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
		fmt.Println("url Scheme", u.Scheme)
		fmt.Println("url Host", u.Host)
		u = buildUrl(u)
		verb, _ := cmd.Flags().GetString("verb")
		payload, _ := cmd.Flags().GetString("data")
		header, _ := cmd.Flags().GetString("header")
		response, err := client.Do(u.String(), verb, payload, header)
		if err != nil {
			log.Println("error making request", err)
			return err
		}
		cout.Header(u.String(), verb)
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
