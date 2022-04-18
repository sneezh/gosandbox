package main

import "math"

func maxConsecutiveElements(str string) int {
	result, currLetterIdx := 0, 0
	for currLetterIdx < len(str) {
		nextLetterIdx := currLetterIdx
		for nextLetterIdx < len(str) && str[nextLetterIdx] == str[currLetterIdx] {
			nextLetterIdx += 1
		}
		result = int(math.Max(float64(result), float64(nextLetterIdx-currLetterIdx)))
		currLetterIdx = nextLetterIdx
	}
	return result
}
