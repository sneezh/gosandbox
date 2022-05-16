package main

func isValid(s string) bool {
	var stack []byte
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if len(stack) == 0 || stack[len(stack)-1] != byte(c) {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
