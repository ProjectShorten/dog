package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// get url from user
	url := flag.String("url", "https://example.com/", "URL to fetch")

	flag.Parse()
	// validate url
	if url == nil {
		fmt.Println("No URL provided")
		return
	}

	if *url == "" {
		fmt.Println("Invalid URL provided")
		return
	}

	// fetch data
	res, err := http.Get(*url)
	if err != nil {
		fmt.Println("Error whilst fetching: ", err)
		return
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to ready body: ", err)

	}

	fmt.Println(string(content))

	// output
}

func fetch() {

}
