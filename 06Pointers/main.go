package main

import "fmt"

func main() {
	fmt.Println("this is go lang class of pointers")
	var pointer *int
	fmt.Println("this is vlaue of pointer", pointer)

	val := 11
	fmt.Println("this is value", val)
	val2 := &val
	fmt.Println("this is value 2", val2)
	val = val + 22
	fmt.Println("this is value+22", val)
	fmt.Println("this is value 2 again", val2)

	fmt.Print("this is refference", &val)
	fmt.Println("this is pointer", *val2)
	fmt.Printf("this is type of val2 %T \n", &val2)
	fmt.Printf("this is type of val %T  \n", val)

}
