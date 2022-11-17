// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

// Входные данные

// Вводится натуральное число.

// Выходные данные

// Выведите ответ на задачу.
package main

import "fmt"

func main() {
	var a int

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}

	var digit int = 1
	for digit <= a {
		fmt.Printf("%v ", digit)
		digit = digit*2
	}
}
