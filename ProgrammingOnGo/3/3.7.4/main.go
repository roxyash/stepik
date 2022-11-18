// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.

// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
// а затем вывести на стандартный вывод в том же формате.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(string(data))
	}
	data = strings.TrimSuffix(data, "\n")
	date, err := time.Parse("2006-01-02 15:04:05", string(data))
	if err != nil {
		panic(err)
	}

	if date.Hour() > 13 {
		fmt.Print(date.Add(time.Second * 86400).Format("2006-01-02 15:04:05"))
	} else {
		fmt.Print(date.Format("2006-01-02 15:04:05"))
	}
}
