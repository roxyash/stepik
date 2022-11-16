// Для решения данной задачи вам понадобится пакет strconv,
// возможно использовать пакеты strings или encoding/csv,
// или даже bufio - вы не ограничены в выборе способа решения задачи.
// В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

// В привычных нам редакторах электронных таблиц присутствует удобное представление числа
// с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой.
// Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

// На стандартный ввод вы получаете 2 таких вещественных числа,
// в качестве результата требуется вывести частное от деления первого числа на второе с точностью
// до четырех знаков после "запятой"
// (на самом деле после точки, результат не требуется приводить к исходному формату).

// P.S. небольшое отступление, связанное с чтением из стандартного ввода.
// Кто-то захочет использовать для этого пакет bufio.Reader.
// Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'),
// то получаете ошибку EOF, а точнее (io.EOF - End Of File). На самом деле это не ошибка,
// а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца.
// Чтобы ошибка была обработана правильно, вы можете поступить так:
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Scan err: %v", err)
	}
	afterComa := strings.Split(text, ";")
	var floatNumbers []float64
	for _, value := range afterComa {
		value2, err := strconv.ParseFloat(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(value, " ", ""), ",", ".")), 64)
		if err != nil {
			fmt.Printf("scrconv.ParseFloat err: %v", err)
		}

		floatNumbers = append(floatNumbers, value2)
	}

	fmt.Printf("%.4f", floatNumbers[0]/floatNumbers[1])
}
