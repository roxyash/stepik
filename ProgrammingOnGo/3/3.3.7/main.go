// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
// которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
// Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
package main //Remove this before push stepik

import ( //Remove this before push stepik
	"fmt"     //Remove this before push stepik
	"strconv" //Remove this before push stepik
) //Remove this before push stepik

func main() { //Remove this before push stepik
	fn := func(a uint) uint {
		if a == 0 {
			return 100
		}
		var str string
		var arr []uint
		for true {
			if a/10 == 0 {
				if (a%10)%2 == 0 && a%10 != 0 {
					arr = append(arr, a%10)
				}
				break
			}
			if (a%10)%2 == 0 && a%10 != 0 {
				arr = append(arr, a%10)
			}
			a = a / 10
		}

		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		for _, value := range arr {
			str += fmt.Sprintf("%v", value)
		}

		value, _ := strconv.ParseUint(str, 10, 0)
		if value == 0 {
			return 100
		}
		return uint(value)
	}

	fmt.Println(fn(0)) //Remove this before push stepik
} //Remove this before push stepik
