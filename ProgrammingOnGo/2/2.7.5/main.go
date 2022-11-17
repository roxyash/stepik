// На вход подается целое число. Необходимо возвести
//  в квадрат каждую цифру числа и вывести получившееся число.

// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
// Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

package main

import "fmt"

func getDigitsFromNumber(digit int64) []int64 {
	var arr []int64
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

func sqrtNumber(arr []int64 ) string{
	var newArr []int64
	for _, value := range arr {
		newArr = append(newArr, value*value)
	}

	for i, j := 0, len(newArr)-1; i < j; i, j = i+1, j-1 {
		newArr[i], newArr[j] = newArr[j], newArr[i]
	}
	
	var str string

	for _, value := range newArr{
		str += fmt.Sprintf("%v", value)
	}

	return str
}

func main(){
	var a int64
	_, err:= fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err)
	}

	digits := getDigitsFromNumber(a)


	fmt.Print(sqrtNumber(digits))
}