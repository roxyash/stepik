// Из натурального числа удалить заданную цифру.

// Входные данные

// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные

// Вывести число без заданных цифр.
package main

import "fmt"

func reverseArr(arr []int) []int {
	var newArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	 }
	return newArr
}

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
		fmt.Printf("%v", arr[i])
	}
}


func removeDigit(arr []int, n int) []int {
	var newArr []int

	for _, value := range arr {
		if value == n {
			continue
		}
		newArr = append(newArr, value)
	}

	return newArr
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}

	digits := reverseArr(getDigitsFromNumber(a))


	printArr(reverseArr(removeDigit(digits, b)))

}
