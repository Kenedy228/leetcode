package task53

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = getMax(nums[i], nums[i]+currentSum)
		maxSum = getMax(maxSum, currentSum)
	}

	return maxSum
}

func getMax(x, y int) int {
	if x > y {
		return x
	}

	return y
}
