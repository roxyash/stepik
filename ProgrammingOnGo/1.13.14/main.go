// Двоичная запись
// Дано натуральное число N. Выведите его представление в двоичном виде.

// Входные данные

// Задано единственное число N

// Выходные данные

// Необходимо вывести требуемое представление числа N.

package main

import "fmt"

func createNumberFromDigit(arr []int) string {
	var str string 
	for i:=len(arr)-1; i>=0; i--{
		str += fmt.Sprintf("%v", arr[i])
	}
	return str
}

func getBinaryFromNumber(a int){ 
	var arr []int
	for true {
		arr = append(arr, a%2)
		a = a/2
		if a == 0 {
			break
		}
	}

	fmt.Println(createNumberFromDigit(arr))

}

func main() {
	var a int 

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}

	getBinaryFromNumber(a)


}