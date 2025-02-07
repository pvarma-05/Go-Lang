package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please Give Rating for Pradeep's Work : ")
	ip, _ := reader.ReadString('\n') // Comma err : input,error (can handle both)
	fmt.Println("Thanks For Rating : ", ip)
	fmt.Printf("This is of type : %T", ip)
}
