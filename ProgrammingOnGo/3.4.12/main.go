// Пришло время для задач, где вы сможете применить полученные знания на практике.

// Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(),
//  которая возвращает 3 значения типа пустой интерфейс. Эта функция использует пакеты encoding/json,
//  fmt, и os - не удаляйте их из импорта. Скорее всего, вам понадобится пакет "fmt",
//  но вы можете использовать любой другой пакет для записи в стандартный вывод не удаляя fmt.

// Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64.
//  Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная математическая операция).
//  Но такие идеальные случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем значении может не относится
// к одной из указанных математических операций.

// Результат выполнения программы должен быть таков:

// в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции с точностью до 4 цифры после запятой (fmt.Printf(%.4f));
// если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке вида: value=полученное_значение: тип_значения (например: value=true: bool)
// если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям, сообщение об ошибке должно иметь вид: неизвестная операция
// Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения увидели, что оно содержит ошибку - печатайте сообщение об ошибке
//  и завершайте работу программы, проверка остальных аргументов значения уже не имеет, а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.

// Удачи!

package main //Remove this string before push stepik

import (//Remove this string before push stepik
	"fmt"//Remove this string before push stepik
)//Remove this string before push stepik

func readTask() (value1, value2, operation interface{}) {//Remove this string before push stepik

	return 5.0, 50, "/" //тут играемся с параметрами//Remove this string before push stepik

}//Remove this string before push stepik

func main() {//Remove this string before push stepik
	value1, value2, operation := readTask()//Remove this string before push stepik
	arr := [...]interface{}{value1, value2, operation}

	var arrNumbers []float64
	mapS := make(map[string]string)

	for _, value := range arr {
		switch value.(type) {
		case float64:
			arrNumbers = append(arrNumbers, value.(float64))
		case string:
			mapS[value.(string)] = value.(string)
		default:
			fmt.Printf("value=%v: %T", value, value)
			return
		}
	}

	var sum float64
	for k := range mapS {
		switch k {
		case "+":
			sum = arrNumbers[0] + arrNumbers[1]
		case "-":
			sum = arrNumbers[0] - arrNumbers[1]
		case "/":
			sum = arrNumbers[0] / arrNumbers[1]
		case "*":
			sum = arrNumbers[0] * arrNumbers[1]
		}
	}

	fmt.Printf("%.4f", sum)

}//Remove this string before push stepik
