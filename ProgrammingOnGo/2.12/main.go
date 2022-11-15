// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита.
//  На вход подается строка-пароль.
// Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}
	
	passwordRune := []rune(a)
	var cntDigit int 
	var cntLat int 
	if len(passwordRune) >=5{
		for _, value := range passwordRune{
			if unicode.IsDigit(value){
				cntDigit++
			}
			if unicode.Is(unicode.Latin, value) {
				cntLat++
			}
		}
		if cntDigit+cntLat == len(passwordRune) {
			fmt.Println("Ok")
		}else{
			fmt.Println("Wrong password")
		}

	}else{
		fmt.Println("Wrong password")
	}
}
