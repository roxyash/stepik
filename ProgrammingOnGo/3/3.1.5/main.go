// Внутри функции main (объявлять функцию не нужно) необходимо написать программу:

// На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться).
//  Для чтения из стандартного ввода импортирован пакет fmt.

// Вам необходимо вычислить результат выполнения функции work для каждого из полученных чисел.
//  Функция work имеет следующий вид:

// func work(x int) int
// Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.

// Однако работа функции work занимает слишком много времени.
//  Выделенного вам времени выполнения не хватит на последовательную обработку каждого числа,
// поэтому необходимо реализовать кэширование уже готовых результатов и использовать их в работе.

// После завершения работы программы результат выполнения будет дополнен информацией о
//  соблюдении установленного лимита времени выполнения.
package main //Delete this str before push stepik

import "fmt" //Delete this str before push stepik

func work(x int) int { //Delete this str before push stepik
	return 1 //Delete this str before push stepik
} //Delete this str before push stepik

func main() { //Delete this str before push stepik
	mapA := make(map[int]int)
	var a int
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Printf("Scan err: %v", err)
		}
		if value, ok := mapA[a]; ok {
			fmt.Printf("%v ", value)
		} else {
			mapA[a] = work(a)
			fmt.Printf("%v ", mapA[a])
		}
	}
} //Delete this str before push stepik
