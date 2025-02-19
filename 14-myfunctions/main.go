package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	displayName()
	res1 := noobadder(5, 7)
	fmt.Println("The Result From Noob Adder is : ", res1)
	res2 := proadder(5, 19, 27, 34, 13, 14, 3, 5, 6, 52, 47, 6, 5, 4, 5, 75, 65, 5)
	fmt.Println("Result of Pro Adder : ", res2)
	res3, mssg := conclusion()
	fmt.Println("My Lucky Number is ", res3)
	fmt.Println(mssg)
}

func noobadder(val1 int, val2 int) int {
	return val1 + val2
}

func proadder(values ...int) int { // can take multiple values
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func displayName() { // Function should not be declared in another function
	fmt.Println("Hello My Name is Pradeep Varma")
}

func conclusion() (int, string) { // returning multiple values
	return 5, "Thank You"
}
