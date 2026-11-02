func findMin(nums []int) int {
	left, right := 0, len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		fmt.Println("left:", left, "mid:", mid, "right:", right)

		if (nums[mid] > nums[right]) {  // min should be at rhs
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
