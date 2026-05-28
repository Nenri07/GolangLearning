package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is Method class")

	me := User{"Umer", "umarbashir@gmail.com", 21}

	name := bufio.NewReader(os.Stdin)
	chnageName, _ := name.ReadString('\n')

	fmt.Println("this is all value of me", me)

	fmt.Printf("this is key value struct %+v \n", me)

	fmt.Printf("this is the name %v, and this is id %v", me.Name, me.Id)
	me.GetName()
	me.ChnageName(chnageName)

}

type User struct {
	Name  string
	Email string
	Id    int
}

func (u User) GetName() {

	fmt.Println("this is name", u.Name)
}

func (u User) ChnageName(name string) {
	fmt.Printf("Name chnaged successfully from %v to %v\n", u.Name, name)
	u.Name = name

}
