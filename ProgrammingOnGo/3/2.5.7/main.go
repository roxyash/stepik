	// На вход подается строка. Нужно определить,
	// является ли она правильной или нет.
	// Правильная строка начинается с заглавной буквы и заканчивается точкой.
	// Если строка правильная - вывести Right иначе - вывести Wrong

	// Маленькая подсказка: fmt.Scan() считывает строку до первого пробела,
	// вы можете считать полностью строку за раз с помощью bufio:

	// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	package main

	import (
		"bufio"
		"fmt"
		"os"
		"unicode"
	)

	func main() {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		textInRune := []rune(text)
		if unicode.IsUpper(textInRune[0]) && textInRune[len(textInRune)-3] == 46 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	}
