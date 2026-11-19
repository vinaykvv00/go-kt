package main

import (
	"fmt"
	"net/url"
)

const baseURL = "https://example.com/"

func GetURL() string {
	return baseURL
}

func main() {
	println("URL:", GetURL())
	println("URL length:", len(GetURL()))

	//parsing
	result, err := url.Parse(GetURL())
	if err != nil {
		println("Error parsing URL:", err.Error())
		return
	}
	println("Parsed URL:", result.String())

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("Query Parameters:", qparams)
	fmt.Println("Type of query parameters:", fmt.Sprintf("%T", qparams))

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: url.Values{"user": {"hitesh"}}.Encode(),
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
