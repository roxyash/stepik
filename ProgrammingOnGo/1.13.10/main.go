// Самое большое число, кратное 7
// Найдите самое большее число на отрезке от a до b, кратное 7 .

// Входные данные
// Вводится два целых числа a и b (a≤b).

// Выходные данные
// Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.
package main

import (
	"fmt"
)

func getAllNumbers(a, b int) []int{ 
	var arr []int
	if a == b {
		arr = append(arr, a)
		return arr 
	} 
	if a<b {
		for i := a; i <=b; i++ {
			arr = append(arr, i)
		}
	}
	if a>b {
		for i := b; i <=a; i++ {
			arr = append(arr, i)
		}
	}

	return arr
}

func getMaxInArray(arr []int) int{
	max := arr[0]
	for _, value := range arr {
		if value%7 == 0 {
			if value > max {
				max = value
			}
		}

	}

	return max
}

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}

	arr := getAllNumbers(a,b)

	if max := getMaxInArray(arr); max%7!=0 {
		fmt.Println("NO")
	}else{
		fmt.Println(max)
	}
}
