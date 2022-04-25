package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{1, 2, 44, 3, 2, 6, 43, 22}, 2))
}

func findKthLargest(nums []int, k int) int {
	return kSelect(nums, 0, len(nums)-1, k)
}

func kSelect(nums []int, left int, right int, k int) int {
	if left == right {
		return nums[left]
	}

	p := right
	l := left
	r := right - 1
	for l <= r {
		for l <= r && nums[l] >= nums[p] {
			l++
		}

		for l <= r && nums[r] < nums[p] {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	p = l

	if p > k-1 {
		return kSelect(nums, left, p-1, k)
	} else if p < k-1 {
		return kSelect(nums, p+1, right, k)
	}

	return nums[p]
}
