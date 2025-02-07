package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	favFooList := []string{"Chicken-Shawarma", "Chicken-Popcorn"} // this declaration is a Slice
	fmt.Println("My Favourite Food List is : ", favFooList)
	fmt.Println("Total Fav Foods : ", len(favFooList)) // we get original length here
	favFooList = append(favFooList, "Kaju-Katli", "Chicken-Loaded")
	fmt.Println("Updated Fav Food List is : ", favFooList)
	fmt.Println("Total Fav Food list now is: ", len(favFooList))

	fmt.Printf("\n")
	favFooList = append(favFooList[:2]) //ignore error
	fmt.Println("Top 2 Fav Foods are : ", favFooList)

	fmt.Printf("\n")

	score := make([]int, 4)
	score[0] = 65
	score[1] = 67
	score[2] = 62
	score[3] = 70

	fmt.Println("Scores are : ", score)
	fmt.Println("Len of Scores :", len(score))

	score = append(score, 80, 90)
	fmt.Println("Updated Scores List : ", score)
	fmt.Println("Len of Updated Scores: ", len(score))

	fmt.Println("Scores are sorted? ", sort.IntsAreSorted(score))
	sort.Ints(score)
	fmt.Println("Sorted Scores :", score)
	fmt.Println("Now Scores are sorted? :", sort.IntsAreSorted(score))

	// remove a value from Slices [Updated Commit]
	fmt.Printf("\n\n")
	courses := []string{"CD", "CNS", "AI", "WAD", "RSM", "CS"}
	fmt.Println("Course List: ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated Courses List: ", courses)

}
