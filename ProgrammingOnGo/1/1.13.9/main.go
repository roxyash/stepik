// Цифровой корень
// Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

// Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .

// По данному числу определите его цифровой корень.

// Входные данные

// Вводится одно натуральное число n, не превышающее 10^710
// 7
//  .

// Выходные данные
// Вывести цифровой корень числа n.
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
		digit = digit/10
	}
	if sumArr(arr) > 9 {
		arr = getDigitsFromNumber(sumArr(arr))
	}
	return arr 
}

func sumArr(arr []int) int {
	var sum int 
	for _, value := range arr {
		sum += value
	}

	return sum
}

func main() {
	var a int

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Err scan: %v", err.Error())
	}

	digits := getDigitsFromNumber(a)
	sum := sumArr(digits)

	fmt.Println(sum)
}
