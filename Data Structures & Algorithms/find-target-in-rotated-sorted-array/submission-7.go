func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		if nums[left] < nums[mid] { // left sorted
			if nums[left] < target && target < nums[mid] { // target in range
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //right sorted
			if nums[mid] < target && target < nums[right] { // target in range
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
