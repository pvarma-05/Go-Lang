package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")
	myTech := make(map[string]string)
	myTech["JS"] = "Javascript"
	myTech["NJS"] = "Node JS"
	myTech["TS"] = "Typescript"
	myTech["RJS"] = "React JS"

	fmt.Println("List of My Tech Stack : ", myTech)
	fmt.Println("NJS Stands for: ", myTech["NJS"])

	delete(myTech, "NJS")
	fmt.Println("List of Updated Tech Stack : ", myTech)
}
