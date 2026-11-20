package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const verbBaseURL = "http://127.0.0.1:3000"

func main() {
	fmt.Println("welcome to verb go web")
	client := &http.Client{Timeout: 5 * time.Second}
	if err := run(client); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(client *http.Client) error {
	fmt.Println("Performing GET request...")
	response, err := client.Get(verbBaseURL + "/")
	if err != nil {
		return fmt.Errorf("GET /: %w (start the demo server with go run ./verbserver)", err)
	}
	if err := printResponse("GET /", response); err != nil {
		return err
	}

	payload := strings.NewReader(`{"coursename":"Let's go with golang","price":0,"platform":"learnCodeOnline.in"}`)
	response, err = client.Post(verbBaseURL+"/post", "application/json", payload)
	if err != nil {
		return fmt.Errorf("POST /post: %w", err)
	}
	if err := printResponse("POST /post", response); err != nil {
		return err
	}

	data := url.Values{
		"firstname": {"hitesh"},
		"lastname":  {"choudhary"},
		"email":     {"hitesh@go.dev"},
	}
	response, err = client.PostForm(verbBaseURL+"/postform", data)
	if err != nil {
		return fmt.Errorf("POST /postform: %w", err)
	}
	return printResponse("POST /postform", response)
}

func printResponse(request string, response *http.Response) error {
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("%s: unexpected HTTP status %s", request, response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("%s: read response: %w", request, err)
	}
	fmt.Printf("%s: %s\n%s\n", request, response.Status, strings.TrimRight(string(body), "\r\n"))
	return nil
}
