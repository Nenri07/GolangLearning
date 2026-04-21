package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Input class"
	fmt.Println(welcome)

	read := bufio.NewReader(os.Stdin)
	fmt.Println("enter my skill on golang")
	//comma ok syntax|| error syntax
	input, _ := read.ReadString('\n')

	fmt.Println("this is rating of my skill", input)
	fmt.Printf("this is type of input %T \n", input)

}
