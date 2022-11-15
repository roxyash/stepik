// Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
// но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка",
// если ошибки нет - результат функции. Функция divide(a int, b int) (int, error)
//  получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error).
//  Пакет main уже объявлен, а нужные пакеты уже импортированы!
package main

import "fmt"

func divide(a,b int) (int, error) {
	return a/b, nil
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return 
	}
	res, err := divide(a, b)
	if err != nil {
        fmt.Println("ошибка")
		return 
    }else{
        fmt.Println(res)
    }

}