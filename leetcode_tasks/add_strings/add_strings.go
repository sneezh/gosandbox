package main

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := byte(0)
	l1 := len(num1) - 1
	l2 := len(num2) - 1

	for l1 >= 0 || l2 >= 0 {
		sum := byte(0)
		if l1 >= 0 {
			sum += num1[l1] - '0'
		}
		if l2 >= 0 {
			sum += num2[l2] - '0'
		}

		sum += carry
		carry = sum / 10
		n := sum % 10
		res = string(n+'0') + res
		l1--
		l2--
	}

	if carry > 0 {
		res = string(carry+'0') + res
	}
	return res
}
