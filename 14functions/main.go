package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("this is functions class")

	inp := bufio.NewReader(os.Stdin)
	inval, err := inp.ReadString('\n')

	if err != nil {
		fmt.Println("this is value of err", err)

	}

	result := chatgpt(inval)

	fmt.Println("processing.... is t you meant ", result)

	addslice, msg := adder(3999, 88, 9, 11, 23)

	fmt.Println("This is total", addslice, "and message\n", msg)

}

func adder(values ...int) (int, string) {
	result := 0
	for _, value := range values {
		result += value
	}
	return result, "this is value"

}

func chatgpt(value string) string {
	result := ""

	if strings.Contains(value, "hello") {
		result := "hey its me goku"
		return result

	} else if strings.Contains(value, "bye") {
		result := "no its nenrei"
		return result
	}
	return result

}
