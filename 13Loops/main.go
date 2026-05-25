package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is class of the loops")

	names := []string{"Umer", "Tayyab", "Ustad g"}

	for n := 0; n < len(names); n++ {
		fmt.Println(names[n])
	}

	for i := range names {
		fmt.Println(names[i])
	}

	for i, value := range names {
		fmt.Printf("this is value of index %v and this is value of the index %v \n", i, value)
	}

	checkvalu := 1
jumper:
	fmt.Println("Yoho ia m here jumped from the index out")

	for checkvalu < 10 {
		if checkvalu == 2 {
			break
		} else if checkvalu == 7 {
			fmt.Println("Jumpping....")
			goto jumper

		} else if checkvalu == 9 {
			checkvalu++
			continue
		}
		fmt.Println("this is index", checkvalu)
		checkvalu++

	}

}
