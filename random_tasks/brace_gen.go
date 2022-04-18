package main

import "fmt"

func main() {
	accum := &[]string{}
	fmt.Println(generate(2, 2, accum))
}

func generate(leftOpen int, leftClosed int, accum *[]string) []string {
	if leftOpen == 0 && leftClosed == 0 {
		return *accum
	}

	if leftOpen > 0 {
		*accum = append(*accum, "(")
		generate(leftOpen-1, leftClosed, accum)
	}
	if leftClosed > leftOpen {
		*accum = append(*accum, ")")
		generate(leftOpen, leftClosed-1, accum)
	}
	return *accum
}
