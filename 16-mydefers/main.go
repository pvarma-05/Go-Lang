package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defers")

	// My POV : defer is a keyword which helps the statement to execute last.
	// Which also Uses stack to store which one to execute first. i.e LIFO Principle
	// Example :

	fmt.Println("Hello Boiss")
	defer fmt.Println("My Name is Pradeep")
	defer fmt.Println("Student @ GITAM University")
	defer fmt.Println("My YOB is 2005") // run the code to know exactly
	// basically first We get Hello Boiss in Output
	// after that in stack we store defer statements - [Name, Place, Year]
	// As it works on LIFO, next the year would be Executed after that place and Name...

	mydefer() // now similarly, we get the output as reverse like - 4 3 2 1 0
	// and also note that function is not defer so it executes after hello boiss statement
}

func mydefer() {
	for i := range 5 {
		defer fmt.Println(i)
	}
}
