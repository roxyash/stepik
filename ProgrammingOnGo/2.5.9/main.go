// Даются две строки X и S.
// Нужно найти и вывести первое вхождение подстроки S в строке X.
// Если подстроки S нет в строке X - вывести -1

package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}
	fmt.Println(strings.Index(a, b))
}
