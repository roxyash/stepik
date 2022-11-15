// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит
// только арабские цифры.

// Выходные данные

// Выведите максимальную цифру, которая встречается во введенной строке.
package main

import "fmt"

func max(runes []rune) int {
	var max int
	for _, value := range runes {
		if int(value) > max {
			max = int(value)
		}
	}

	return max
}

func main() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err)
	}

	arrRune := []rune(a)
	fmt.Print(string(rune(max(arrRune))))
}
