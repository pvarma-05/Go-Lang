package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int
	fmt.Println("The Default Value of Pointer is : ", ptr)

	luckyNum := 05
	myptr := &luckyNum
	fmt.Println("My Lucky Number is : ", *myptr)
	fmt.Println("My Lucky Number Stored in :", myptr)

	*myptr = *myptr + 4
	fmt.Println("New Lucky Number is : ", luckyNum) // the value gets updated
}
