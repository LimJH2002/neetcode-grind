func productExceptSelf(nums []int) []int {
	leftPass := make([]int, len(nums))
	rightPass := make([]int, len(nums))
	product := make([]int, len(nums))

	i, j := 0, len(nums) - 1
	leftRunning, rightRunning := 1, 1
	for i < len(nums) && j >= 0 {
		leftPass[i], rightPass[j] = leftRunning, rightRunning
		leftRunning *= nums[i]
		rightRunning *= nums[j]
		i++
		j--
	}

	for i := 0; i < len(nums); i++ {
		product[i] = leftPass[i] * rightPass[i]
	}

	return product
}
