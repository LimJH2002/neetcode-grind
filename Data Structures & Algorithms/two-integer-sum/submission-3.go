func twoSum(nums []int, target int) []int {
    directory := make(map[int] int)

	for i := 0; i < len(nums); i++ {
		candidate, ok := directory[target - nums[i]]
		if ok { // valid candidates
			return []int{candidate, i}
		}
		directory[nums[i]] = i
	}

	return []int{}
}
