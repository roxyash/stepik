package main

import "fmt"

// Напишите функцию которая принимает канал и строку.
// Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
func task(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func main() {
	c := make(chan string)
	go task(c, "str")

	fmt.Println(<-c)
}
