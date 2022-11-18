// На стандартный ввод подается строковое представление даты и времени в следующем формате:

// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

// Wed Apr 16 05:20:00 +0600 1986

// Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.

package main

import (
	"fmt"
	"time"
)

func main() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}

	date, err := time.Parse(time.RFC3339, a)
	if err != nil {
		panic(err)
	}

	newDate := date.Format(time.UnixDate)

	fmt.Println(newDate)
}
