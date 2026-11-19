package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com/"

func main() {
	println("Hello, World! web request")
	println("URL:", url)

	response, err := http.Get(url)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	defer response.Body.Close()
	println("Response Status:", response.Status)
	fmt.Printf("Response type is: %T\n", response)

	// You can read the response body here if needed
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println("Error reading body:", err.Error())
		return
	}
	fmt.Println("Response Body is in bytes:", body)
	fmt.Println("Response Body:", string(body))
}
