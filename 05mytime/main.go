package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("we are talking about time")

	presentTime := time.Now()

	fmt.Println("present time and date is ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2002, time.June, 7, 5, 45, 12, 2, time.Local)
	fmt.Println("this is created date", createdDate)

	fmt.Println("this is formated date and time", createdDate.Format("01-02-2006 Monday"))
}
