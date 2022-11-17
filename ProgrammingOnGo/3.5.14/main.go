package main

import (
	"fmt"
	"os"
	"strings"
)

func returnIndex(str []string) int {
	for index, value := range str {
		if value == "0" {
			return index + 1
		}
	}

	return 0
}
func main() {
	file, err := os.ReadFile("txt.txt")
	if err != nil {
		fmt.Println(err)
	}

	str := strings.Split(string(file), ";")
	fmt.Print(returnIndex(str))
}
