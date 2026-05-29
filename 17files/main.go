package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("This is files class")

	content := "this is file new content"

	file, err := os.Create("./checkOS")
	checkNillError(err)
	fmt.Println("file is", file)
	length, err := io.WriteString(file, content)
	checkNillError(err)

	fmt.Println("length after write in file", length)

	defer file.Close()
	readFile("./checkOS")
}

func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	/// ioutil deprecated now os.readfile
	datainbyte, err := os.ReadFile(filename)
	checkNillError(err)

	fmt.Println("this is byte data", string(datainbyte))

}
