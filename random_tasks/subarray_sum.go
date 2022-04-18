package main

func subarraySum(arr []int, target int) bool {
	right, currentSum := 0, 0
	for left := range arr {
		if left > 0 {
			currentSum -= arr[left-1]
		}

		for right < len(arr) && currentSum < target {
			currentSum += arr[right]
			right++
		}

		if currentSum == target {
			return true
		}
	}

	return false
}
