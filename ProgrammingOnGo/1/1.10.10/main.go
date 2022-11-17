// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

// Входные данные

// Программа получает на вход два числа. Гарантируется,
// что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

// Выходные данные

// Программа должна вывести цифры, которые имеются в обоих числах,
// через пробел. Цифры выводятся в порядке их нахождения в первом числе.

// Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

package main

import "fmt"

func getDigitsFromNumber(digit int) []int {
	var arr []int
	for true {
		if digit/10 == 0 {
			arr = append(arr, digit)
			break
		}

		arr = append(arr, digit%10)
		digit = digit / 10
	}
	return arr
}

func printArr(arr []int) {
	for i := len(arr)-1; i >=0; i-- {
		fmt.Printf("%v ", arr[i])
	}
}

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Error scan: %v", err.Error())
	}

	var digits []int
	digitOne := getDigitsFromNumber(a)
	digitTwo := getDigitsFromNumber(b)

	for _, value := range digitOne {
		for _, value2 := range digitTwo {
			if value == value2 {
				digits = append(digits, value)
				continue
			}
		}
	}

	printArr(digits)
}
