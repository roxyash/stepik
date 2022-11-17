// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}

	fmt.Print(math.Sqrt(a*a+b*b))
}
