package main

import "fmt"

var NameofMe string = "umer"

const loggedIn bool = true

func main() {
	var name string = "umer"
	fmt.Println(name)
	fmt.Printf("this is type of name %T \n", name)

	var age uint32 = 22
	fmt.Printf("this is type of age %T \n", age)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("this is type of isLoggedIn %T \n", isLoggedIn)

	var smallFloat float64 = 3.333333
	fmt.Println(smallFloat)
	fmt.Printf("this is type of small float %T \n", smallFloat)

	var checkVeraiable int
	fmt.Println(checkVeraiable)
	fmt.Printf("this is type of verable %T \n", checkVeraiable)

	var checkVeriable2 string
	fmt.Println(checkVeriable2)
	fmt.Printf("this is type of verable %T \n", checkVeriable2)

	var checkVeriable3 bool
	fmt.Println(checkVeriable3)
	fmt.Printf("this is type of verable %T \n", checkVeriable3)
	// wolerace operator
	newName := "umer"
	fmt.Println(newName)

	fmt.Println(NameofMe)
	fmt.Println(loggedIn)

}
