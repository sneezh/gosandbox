package arrays

func sortedSquares(nums []int) []int {
	var i, j = 0, len(nums) - 1
	arr := make([]int, j+1)
	for z := j; z >= 0; z-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			arr[z] = nums[i] * nums[i]
			i++
		} else {
			arr[z] = nums[j] * nums[j]
			j--
		}
	}
	return arr
}
