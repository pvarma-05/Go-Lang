package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Different Web Requests in GO")

	fmt.Println("In this Code we send different Requests to the Dummy Server")

	fmt.Println("Select which type of Request you want to send http://localhost:3000")
	fmt.Printf("\n1. GET\n2. POST\n3. POSTFORM\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input : ")
	ip, _ := reader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(ip), 32, 32)

	switch n {
	case 1:
		performGetRequest()
	case 2:
		performPostJsonRequest()
	case 3:
		performPostFormRequest()
	default:
		fmt.Println("Enter 1,2,3 numbers")
	}
}

func performGetRequest() { // this performs GET Request
	const url = "http://localhost:3000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	bytes, _ := responseString.Write(content)

	fmt.Println("Byte Count:", bytes)
	fmt.Println("Message for Get Request is : ")
	fmt.Println(responseString.String())
}

func performPostJsonRequest() { // this performs GET Request
	const url = "http://localhost:3000/post"

	var reqBody = strings.NewReader(`

		{
			"course" : "GO Lang",
			"price"	: 0,
			"platform" : "youtube"
		}

	`)

	response, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Message for Post Request:")
	fmt.Println(string(content))

}
func performPostFormRequest() { //This Performs PostForm Reqest
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("Country", "India")
	data.Add("State", "Andhra Pradesh")
	data.Add("City", "Gajuwaka")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println("Message is : ")
	fmt.Println(string(content))
}
