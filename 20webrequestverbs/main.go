package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is another request class with our own server in node")
	GetData()
}

func GetData() {
	const url = "http://localhost:8000/get"
	response, err := http.Get(url)
	CheckNillError(err)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder

	datasize, err := responseString.Write(content)
	CheckNillError(err)
	fmt.Println("this is bytedata", datasize)

	fmt.Println("this is orignal data", responseString.String())

}

func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}
