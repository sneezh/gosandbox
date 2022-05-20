package main

import "fmt"

func main() {
	a := "Привет, World 1!"

	for _, s := range a {
		fmt.Printf("key: %s\n", s)
	}
}
