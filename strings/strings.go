package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	// 5
	fmt.Println(len(a))

	a = "привет"
	// 12
	fmt.Println(len(a))

	strings.ToUpper(a)
	// 12
	fmt.Println(len(a))

	b := []rune(a)
	b[0] = 'П'

	// utf8 codes
	fmt.Println(b)
	a = string(b)

	// Привет
	fmt.Println(a)
}
