package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Welcome to Dice Game using Switch Case")
	fmt.Printf("\n")
	fmt.Println("Rolling the Dice!!!")
	dice := rand.IntN(6) + 1
	fmt.Println("You Got : ", dice)

	switch dice {
	case 1:
		fmt.Println("Now U can move 1 step ahead")
	case 2:
		fmt.Println("Now U can move 2 steps ahead")
	case 3:
		fmt.Println("Now U can move 3 steps ahead")
	case 4:
		fmt.Println("Now U can move 4 steps ahead")
	case 5:
		fmt.Println("Now U can move 5 steps ahead")
	case 6:
		fmt.Println("Now U can move 6 steps ahead and U get One More Chance to Roll the Dice")
	default:
		fmt.Println("There is Something Wrong !!!")
	}
}
