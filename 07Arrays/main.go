package main

import "fmt"

func main() {
	fmt.Println("this is array section")

	var array [3]int
	array[2] = 3

	fmt.Println("this is array", array)

	var array3 = [3]string{"abc", "ahsan", "ali"}

	fmt.Println("this is type of array", len(array3), "this si array", array3)

	fmt.Printf("this is type of array %T \n", array3)
	fmt.Printf("this is first array type %T \n ", array)
}
