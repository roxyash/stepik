// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
package main

import (
	"fmt"
	"strings"
)

func main(){
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}
	for _, value := range a {
		if strings.Count(a, string(value)) > 1 {
			a = strings.ReplaceAll(a, string(value), "")
		}
	}

	fmt.Println(a)
}