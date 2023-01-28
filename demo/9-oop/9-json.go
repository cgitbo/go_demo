package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Prices int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {

	movie := Movie{
		Title:  "Star Wars",
		Year:   1989,
		Prices: 2000,
		Actors: []string{"George Washington", "John Adams"},
	}

	// 结构体转json
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonStr))

	// json 转结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", myMovie)

}
