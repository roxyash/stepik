// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readText() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		return s
	}

	return ""
}
func main() {
	str := readText()

	datesStr := strings.Split(str, ",")
	var dates []time.Time
	for _, dateStr := range datesStr {
		date, err := time.Parse("02.01.2006 15:04:05", dateStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		dates = append(dates, date)
	}

	for i, j := 0, len(dates)-1; i < j; i, j = i+1, j-1 {
		dates[i], dates[j] = dates[j], dates[i]
	}
	

	var max, min time.Time
	if dates[0].Second() > dates[1].Second() {
		max = dates[0]
		min = dates[1]
	} else {
		max = dates[1]
		min = dates[0]
	}
	dur := max.Sub(min)

	fmt.Println(dur)
}
