package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("this is Slices class")

	var scores = []float64{3.33, 7.44, 2.22}
	fmt.Println("this is scores list", scores)
	fmt.Printf("this is type of list %T \n", scores)

	scores = append(scores, 2.22, 9.99)
	fmt.Println("this is update dlist", scores)

	scores = append(scores[2:])
	fmt.Println("this us 3rd update", scores)

	scores = append(scores[0:2])
	fmt.Println("this is the 4th update", scores)

	checkList := make([]string, 2)
	checkList[0] = "Umer"
	checkList[1] = "Tayyab"
	//checkList[2] = "Ahsan"

	fmt.Println("this is new array or list aka slice", checkList)
	fmt.Printf("this is slice type new %T\n", checkList)

	checkList = append(checkList, "Ahsan", "Ibtisam")
	fmt.Println("new array or list", checkList)

	sort.Strings(checkList)
	fmt.Println("this is sorted array or slice", checkList)

	///remove from slice in specifc index

	var newRemove = []string{"Nextjs", "Python", "Golang"}
	fmt.Println("this is new list2", newRemove)

	//to remove specifc index we use logic

	var index int = 1

	newRemove = append(newRemove[:index], newRemove[index+1:]...)
	fmt.Println("this is removed list", newRemove)

}
