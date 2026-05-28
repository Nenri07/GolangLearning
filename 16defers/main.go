package main

import "fmt"

func main() {
	defer fmt.Println("this is first but defer")
	defer fmt.Println("this is second but defer")
	fmt.Println("This is class for Deferes")
	defer fmt.Println("this is 4th but defer")

	getDefs()
}

func getDefs() {
	for i := 0; i < 7; i++ {
		defer fmt.Println("this is loop", i)
	}
}
