package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to If Else")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any number : ")
	ip, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(ip), 64)

	// 1st way of If-Else
	if err != nil {
		fmt.Println("U have Error in your code : ", err) // error comes if the input is string
	} else if num > 10 {
		fmt.Println("Provided Input is Greater than 10")
	} else if num < 10 {
		fmt.Println("Provided Input is Less than 10")
	} else {
		fmt.Println("Provided Input is equal to 10")
	}

	//2nd Wayfmt.Printf("Enter any number : ")
	fmt.Printf("Enter Another Number : ")
	ip2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseInt(strings.TrimSpace(ip2), 36, 36)

	if err != nil {
		fmt.Println("There is an error in input :", err)
	} else if n := num2; n%2 == 0 {
		fmt.Println("The Provided Input is Even")
	} else {
		fmt.Println("The Provided Input is Odd")
	}

}
