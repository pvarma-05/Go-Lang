package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversions")

	fmt.Printf("Please rate Pradeep's Portfolio from 1 to 5 : ")
	reader := bufio.NewReader(os.Stdin)
	ip, _ := reader.ReadString('\n')

	fmt.Println("Thanks For Rating : ", ip)
	fmt.Printf("Type of above rating : %T\n\n", ip)
	fmt.Println("But Pradeep has rated himself 5 more points")
	numberRating, err := strconv.ParseFloat(strings.TrimSpace(ip), 64)
	if err != nil {
		fmt.Println("There is an error in your input : ", err)
	} else {
		fmt.Println("So that makes his total of ", numberRating+5, " / 10")
		fmt.Printf("Type of added rating is %T", numberRating)
	}
}
