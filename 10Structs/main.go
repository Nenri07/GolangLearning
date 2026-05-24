package main

import "fmt"

func main() {
	fmt.Println("this is struct class")

	me := User{"Umer", "umarbashir@gmail.com", 21}

	fmt.Println("this is all value of me", me)

	fmt.Printf("this is key value struct %+v \n", me)

	fmt.Printf("this is the name %v, and this is id %v", me.Name, me.Id)
}

type User struct {
	Name  string
	Email string
	Id    int
}
