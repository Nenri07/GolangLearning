package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("This is another request class with our own server in node")
	GetData()
	postdata()
	postFormData()
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

func postdata() {
	const url string = "http://localhost:8000/post"

	body := strings.NewReader(`
	{
	"name":"nenrei",
	"lang":"Japnease"
	}
	`)

	response, err := http.Post(url, "application/json", body)
	CheckNillError(err)
	fmt.Println("this is raw response", response)

	defer response.Body.Close()

	bytedata, _ := io.ReadAll(response.Body)

	fmt.Println("this is all formated data", string(bytedata))

}

func postFormData() {
	const murl string = "http://localhost:8000/postform"
	//formdata preparing
	formdata := url.Values{}

	formdata.Add("name", "nenrei")
	formdata.Add("email", "streamsy.dev")
	formdata.Add("age", "22")

	res, err := http.PostForm(murl, formdata)

	CheckNillError(err)
	defer res.Body.Close()

	byteData, _ := io.ReadAll(res.Body)

	fmt.Println("this is formdata with response", string(byteData))

}
