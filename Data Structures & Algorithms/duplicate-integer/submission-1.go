func hasDuplicate(nums []int) bool {
    duplicates := make(map[int] bool)

	for i := 0; i < len(nums); i++ {
		if duplicates[nums[i]] {
			return true
		}
		duplicates[nums[i]] = true
	}

	return false
}
