package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("this is switch class")

	task := []User{}
	id := 0
	fmt.Println("these are just solution", task, id)
loop:
	for {
		fmt.Println(
			`	choose option to process
				add task press 1
				complete task by id press 2
				cancel task by id press 3
				see all task press 4
				see task by priority press 5 
				press 6 to exit`)

		read := bufio.NewReader(os.Stdin)
		option, _ := read.ReadString('\n')

		option = strings.TrimSpace(option)
		switch option {
		case "1":
			fmt.Println("Enter the title")
			title := bufio.NewReader(os.Stdin)
			tVal, _ := title.ReadString('\n')
			fmt.Println("Enter priority")
			priority := bufio.NewReader(os.Stdin)
			ptVal, _ := priority.ReadString('\n')

			ptCon, err := strconv.Atoi(strings.TrimSpace(ptVal))

			if err != nil {
				fmt.Println("this is error", err)
			} else {
				task = append(task, User{Id: id, Title: tVal, Status: "pending", Priority: ptCon})
				id++
				fmt.Printf("task added and this is task no %v\n", id)

			}

		case "2":
			fmt.Println("Enter task id to make it complete")
			complete := bufio.NewReader(os.Stdin)
			completed, _ := complete.ReadString('\n')

			compId, err := strconv.Atoi(strings.TrimSpace(completed))

			if err != nil {
				fmt.Println("this is err", err)
			} else {
				for key, value := range task {
					if value.Id == compId {
						task[key].Status = "completed"
						fmt.Printf("Task completed\n")
						break
					}
				}

			}

		case "3":
			fmt.Println("Enter task id to canecel")
			cancel := bufio.NewReader(os.Stdin)
			canID, _ := cancel.ReadString('\n')

			cId, err := strconv.Atoi(strings.TrimSpace(canID))
			if err != nil {
				fmt.Println("This is error", err)
			} else {
				for i, v := range task {
					if v.Id == cId {
						task[i].Status = "canceled"
						fmt.Printf("TaskCanceled \n")
						break
					}
				}
			}
		case "4":

			fmt.Printf("this is without sort %v+\n", task)
			sort.Slice(task, func(i, j int) bool {
				return task[i].Id < task[j].Id
			})
			fmt.Printf("this is sorted by id %v+\n", task)

		case "5":
			sort.Slice(task, func(i, j int) bool {
				return task[i].Priority < task[j].Priority
			})

			fmt.Printf("this si sort by priority tasks %v+\n", task)
		case "6":
			break loop
		}

	}

}

type User struct {
	Id       int
	Title    string
	Status   string
	Priority int
}
