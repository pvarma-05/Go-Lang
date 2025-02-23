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
	encodeJson() //this func is for encoding JSON
	decodeJson() //this func is for decoding JSON
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

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "React JS",
                "cost": 20,
                "website": "Youtube",
                "tags": [
                        "Frontend",
                        "JS"
                ]
    }
	`)

	var mycourse course

	//1st way

	isValid := json.Valid(jsonDataFromWeb)
	if isValid {
		fmt.Println("JSON Data is Valid")
		json.Unmarshal(jsonDataFromWeb, &mycourse)
		// fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// some cases where we want to add key to data
	// 2nd way

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	// fmt.Printf("%#v\n", myData)

	//more formatted Version
	for k, v := range myData {
		fmt.Printf("Key is %v and Value is %v and Type is %T\n", k, v, v)
	}

}
