package main

const MOD = int(1e9 + 7)

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	dp := make([]int, n, n)
	stack := make([]int, 0, n)
	result := 0
	for i := 0; i != n; i++ {
		for len(stack) != 0 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			prev := stack[len(stack)-1]
			dp[i] = (arr[i]*(i-prev) + dp[prev]) % MOD
		} else {
			dp[i] = (arr[i] * (i + 1)) % MOD
		}
		stack = append(stack, i)
		result = (result + dp[i]) % MOD
	}

	return result
}

func sumSubarrayMinsN2(arr []int) int {
	sum := 0

	for k, v := range arr {
		curMin := v
		sum += v

		if k == len(arr)-1 {
			break
		}

		for _, sv := range arr[k+1:] {
			if curMin > sv {
				curMin = sv
			}

			sum += curMin
		}
	}

	return sum
}
