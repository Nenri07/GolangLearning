package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://dummyjson.com/RESOURCE/?limit=10&skip=5&select=key1&select=key2&select=key3"

func main() {
	fmt.Println("This is URLS class")
	result, _ := url.Parse(myurl)
	fmt.Println("this is url parsing", result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("this is query param type %T\n ", qparams)

	for index, value := range qparams {
		fmt.Printf("this is query index %v and this is param value %v\n", index, value)
	}
	partofurl := &url.URL{
		Scheme:   "https",
		Host:     "dummyjson.com",
		Path:     "/RESOURCE",
		RawQuery: "user=umer",
	}

	anotherUrl := partofurl.String()
	fmt.Println(anotherUrl)
}
