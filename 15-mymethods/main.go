package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")

	pvarma := User{"Pradeep Varma", "spvpenumatsa2@gmail.com", 20, 123456789, true}
	fmt.Println("New user Log : ", pvarma)
	fmt.Printf("Formatted New User's Details : \n%+v\n", pvarma) //placeholder to format details
	fmt.Printf("New User's Name : %v & New User's Email: %v\n", pvarma.Name, pvarma.Email)
	pvarma.getStatus()
}

type User struct {
	Name       string
	Email      string
	Age        int
	Mobile     int
	IsVerified bool
}

func (u User) getStatus() { // This is the Method Declaration
	fmt.Println("Verification Status of the User is ", u.IsVerified)
}
