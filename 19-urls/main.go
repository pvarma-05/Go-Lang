package main

import (
	"fmt"
	"net/url"
)

const URL = "https://tanjore.vercel.app:5000/learn?coursename=nextjs&paymentid=qwr2311awef"

func main() {
	fmt.Println("Welcome to URLs in GO Lang")
	fmt.Println("My URL : ", URL)

	r, _ := url.Parse(URL) // we use this to parse the URL
	fmt.Printf("\n")
	fmt.Println("Scheme is ", r.Scheme)
	fmt.Println("Host is ", r.Host)
	fmt.Println("Path is ", r.Path)
	fmt.Println("Port is :", r.Port())
	fmt.Println("Host is ", r.RawQuery)

	qparams := r.Query()
	fmt.Printf("\nParams are of type: %T\n", qparams)
	fmt.Println("The Param Values of URL are:")
	for i, v := range qparams {
		fmt.Println("Index : ", i, " and Value is: ", v)
	}

	partsUrl := &url.URL{ // this is used to construct an URL
		Scheme:  "https",
		Host:    "tanjore.vercel.app",
		Path:    "/login",
		RawPath: "username=pvarma-05",
	}

	newUrl := partsUrl.String()
	fmt.Println("Our New URL is ", newUrl) //Constructed URL

}
