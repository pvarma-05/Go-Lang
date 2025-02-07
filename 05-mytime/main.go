package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Section")

	currentTime := time.Now() //present Time
	fmt.Println(currentTime)
	fmt.Println("Formatted Time is : ", currentTime.Format("02-01-2006 Monday 15:04:05"))

	myDOB := time.Date(2005, time.April, 23, 17, 30, 0, 0, time.UTC) //past time
	fmt.Println("My DOB is : ", myDOB.Format("02-01-2006 Monday 15:04:05"))
}
