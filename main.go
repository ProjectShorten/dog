package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// get url from user
	url := flag.String("url", "", "URL to fetch")

	flag.Parse()
	// validate url

	// fetch data
	res, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)

	}

	fmt.Println(string(content))

	// output
}