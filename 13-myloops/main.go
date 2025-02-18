package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops :D")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days in a week are : ", days)

	//we use FOR to write any Loop.(no while, and also no pre increment op)

	//1st way of writing Loops
	fmt.Printf("1st Loop Output: \n")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//2nd Way of Writing loops
	fmt.Printf("2nd Loop Output: \n")
	for d := range days {
		fmt.Println(days[d])
	}

	// 3rd way
	fmt.Println("3rd Loop Output:")
	for index, day := range days {
		fmt.Printf("The Index is %v and day is %v\n", index, day)
	}

	// 4th way - While in Other Langs
	fmt.Println("4th Loop Initiated")
	anyValue := 1
	for anyValue < 10 {
		if anyValue == 2 {
			anyValue++
			continue
		}
		if anyValue == 5 {
			goto secret

		}
		if anyValue == 7 {
			break
		}
		fmt.Println("Value is : ", anyValue)
		anyValue++
	}

secret:
	fmt.Println("This is My Secret")

}
