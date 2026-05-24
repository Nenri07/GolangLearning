package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("this is Controll flow class 1")

	var something = []string{"new", "prev", "del"}
	fmt.Println("this is value first time", something)

	sort.Strings(something)

	if n, found := slices.BinarySearch(something, "del"); found {
		something = append(something[:n], something[n+1:]...)
		fmt.Printf("The value is on index %v got deleted\n", n)
	} else {

		something = append(something, "del")
		fmt.Println("the value is added")
	}

	fmt.Println("This is controlled flow slice final", something)
}
