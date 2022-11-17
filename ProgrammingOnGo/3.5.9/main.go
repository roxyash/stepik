package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// strconv, bufio, os, io
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		s := scanner.Text()
		number, _ := strconv.Atoi(s)
		sum += number
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
