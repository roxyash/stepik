// По данному числу n закончите фразу "На лугу пасется..." \
// одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

// Входные данные

// Дано число n (0<n<100).

// Выходные данные

// Программа должна вывести введенное число n и одно из слов (на латинице):
// korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

package main

import "fmt"

func main() {
	var a int

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Error scan: %v", a)
	}

	str := "korov"

	if a <= 10 || a >= 15 {
		switch a % 10 {
		case 1:
			fmt.Printf("%v %v", a, str+"a")
		case 2, 3, 4:
			fmt.Printf("%v %v", a, str+"y")
		case 5, 6, 7, 8, 9, 0:
			fmt.Printf("%v %v", a, str)
		}
	} else {
		switch a {
		case 11, 12, 13, 14:
			fmt.Printf("%v %v", a, str)
		}
	}

}
