package main

import "fmt"

const Location = "Vizag" // if starting letter is Capital then it is Public Variable

func main() {
	fmt.Printf("Welcome to Variables\n\n")

	// normal declaration

	var name string = "Pradeep Varma"
	fmt.Println("My Name is ", name)
	fmt.Printf("This is of type %T\n", name)

	var age int32 = 20
	fmt.Println("My Age is ", age)
	fmt.Printf("This is of type %T\n\n", age)

	// implicit

	var email = "spvpenumatsa2@gmail.com"
	fmt.Println("My Mail id is " + email)
	fmt.Printf("This of type %T\n", email)

	var num = 123456789
	fmt.Println("My Number is ", num)
	fmt.Printf("This is of type %T\n\n", num)

	// no var style
	projectCount := 2
	fmt.Println("I have Completed ", projectCount, " Projects")

	fmt.Println("I Live in ", Location) // Public variable
}
