package client

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func makePostRequest(url string) {
	fmt.Println("making post request")
}


func Do(url string, args ...string) {
	var verb = "GET"
	if len(args) >  0 {
		verb = args[0]
	}
	if verb == "POST" {
		makePostRequest(url)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Can't make request, error: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response, error: ", err)
		return
	}
	
	fmt.Println(resp.Status)
	fmt.Println(string(body))

}