package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("rate the linux")

	//create the reader first with bufio and os packages
	reader := bufio.NewReader(os.Stdin)

	//using reader read from the input

	input, _ := reader.ReadString('\n')

	fmt.Println("this is the overall rating", input)

	//genrally input is like in string and have to use to convert that in the specific type like c#
	converted, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// and to also trim the \n we use new library and function

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("this is new Ratting", converted+5)
	}
}
