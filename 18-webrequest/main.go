package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://tanjore.vercel.app"

func main() {
	fmt.Println("Welcome to Web Requests")
	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of Type : %T\n", response)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Data of the Website:")
	fmt.Println(string(data))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
