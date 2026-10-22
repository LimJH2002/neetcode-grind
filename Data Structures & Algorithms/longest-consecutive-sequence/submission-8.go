func longestConsecutive(nums []int) int {
	numsDir := make(map[int]bool)

	for _, num := range nums { // build the hashmap
		numsDir[num] = true
	}

	longest := 0;
	for num := range numsDir {
		if numsDir[num-1] {
			continue
		}
		length := 1

		for numsDir[num+length] {
			length++
		}
		longest = max(longest, length)
	}

	return longest
}
