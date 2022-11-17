package main

import (
    "fmt" // пакет используется для проверки ответа, не удаляйте его
    "strings"
)

type Stringer interface {
    String() string
}

type Battery struct {
	Text string
}

func (b *Battery) StringTest() string { //Rename name method on String before push stepik 
	var arrStr []string
	for _, value := range b.Text {
		if string(value) == "0"{
			arrStr = append(arrStr, " ")
		}
	}

	for _, value := range b.Text {
		if string(value) == "1"{
			arrStr = append(arrStr, "X")
		}
	}
	
    return "["+strings.Join(arrStr, "")+"]"
}
func main() {
   	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err)
	}

	batteryForTest  := &Battery{Text: a}
	fmt.Println(batteryForTest.StringTest()) //Delete this string before push on stepik 
	
} //Delete this string before push on stepik 
// } Скобка, закрывающая функцию main() вам не видна, но она есть