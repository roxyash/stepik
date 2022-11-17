//  На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
package main

import "fmt"



func main(){
	var a string 
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("Scan err: %v", err.Error())
	}
	var arr []rune
	for index, value := range a {
		if index%2 ==0 {
			continue
		}
		arr = append(arr, value)
	}

	fmt.Println(string(arr))

}