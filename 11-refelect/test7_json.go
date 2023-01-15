package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜劇之王", 2000, 10, []string{"xingye", "karen"}}

	//encoding過程
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//decode jsonstr ->結構體
	//jsonstr = {"title":"喜劇之王","Year":2000,"Price":10,"actors":["xingye","karen"]}
	myMovie := Movie{}
	json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
