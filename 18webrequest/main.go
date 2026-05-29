package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("This is web request class")
	response, err := http.Get(url)
	checkNillError(err)

	fmt.Println("this is response", response)
	fmt.Printf("this is response type %T\n", response)
	defer response.Body.Close()

	bytedata, err := io.ReadAll(response.Body)
	checkNillError(err)
	content := string(bytedata)
	fmt.Println("this is byte data in string", content)
}

func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}
