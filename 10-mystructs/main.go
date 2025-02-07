package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	pvarma := User{"Pradeep Varma", "spvpenumatsa2@gmail.com", 20, 123456789, true}
	fmt.Println("New user Log : ", pvarma)
	fmt.Printf("Formatted New User's Details : \n%+v\n", pvarma)
	fmt.Printf("New User's Name : %v & New User's Email: %v\n", pvarma.Name, pvarma.Email)
}

type User struct {
	Name       string
	Email      string
	Age        int
	Mobile     int
	IsVerified bool
}
