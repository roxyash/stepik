package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ParseStruct struct {
    Id int `json:"global_id"`
}

func main() {
	file, err := os.ReadFile("parse.json")
	if err != nil {
		fmt.Println("Readfile err")
	}

	data := []ParseStruct{}

	err = json.Unmarshal(file, &data)

	var sum int64 
	
	for _, value := range data {
		sum+= int64(value.Id)
	}

	fmt.Println(sum)
}
