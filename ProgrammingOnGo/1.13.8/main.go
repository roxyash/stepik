// Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные

// Вводится натуральное число N, а затем N целых чисел последовательности.

// Выходные данные

// Выведите количество минимальных элементов последовательности.
package main

import "fmt"

func main() {
	var a, b int
	var c []int

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err)
	}

	for i:=0; i< a; i ++ {
		_, err := fmt.Scan(&b)
		if err != nil {
			fmt.Printf("Scan err in loop err: %v", err.Error())
		}
		c = append(c, b)
	}


	min := c[0]
	var minArr []int 
	for _, value := range c {
		if value < min {
			min = value
		}
	}

	for _, value := range c {
		if value == min {
			minArr = append(minArr, value)
		}
	}

	fmt.Println(len(minArr))

}
