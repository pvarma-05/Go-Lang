package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"cost"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Handling")
	encodeJson()
}

func encodeJson() {
	courses := []course{
		{"GO Lang", 20, "Youtube", "abc123", []string{"Backend", "CLI"}},
		{"Jr Software Developer", 60, "Coursera", "def465", []string{"SDE", "Full-Stack", "Java"}},
		{"Full Stack Developer", 50, "Coursera", "deq234", []string{"C#", "Full-Stack", "Copilot"}},
		{"React JS", 20, "Youtube", "abc975", []string{"Frontend", "JS"}},
	}

	res, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", res)
}
