package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	courseList := [7]string{"CD", "CNS", "AI", "WAD", "RSM", "CS"}
	fmt.Println("6th Semester Courses : ", courseList)
	fmt.Println("total subjects are : ", len(courseList)) // original len is 6 but the len shows 7

	favFooList := []string{"Chicken-Shawarma", "Chicken-Popcorn", "Kaju-Katli", "Chicken-Loaded"}
	fmt.Println("My Favourite Food List is : ", favFooList)
	fmt.Println("Total Fav Foods : ", len(favFooList)) // we get original length here
}
