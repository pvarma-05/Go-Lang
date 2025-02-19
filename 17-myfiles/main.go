package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")
	ip := "Hello Guyss, I am Pradeep Varma"

	file, err := os.Create("output.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, ip)
	checkNilErr(err)

	fmt.Println("Length of Given Input is : ", length)
	defer file.Close()
	readFile("./output.txt")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {
	data, err := os.ReadFile(file)
	checkNilErr(err)
	fmt.Println("The Contents of File are :", string(data))
}
