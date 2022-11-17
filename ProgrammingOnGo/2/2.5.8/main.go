// На вход подается строка. Нужно определить, является ли она палиндромом. 
// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
//  (например, "топот", "око", "заказ").

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(arr []rune) string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	reverseArr := []rune(text)
	arr := []rune(text)

	if strings.TrimSpace(string(arr)) == strings.TrimSpace(reverseString(reverseArr)) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}



}
