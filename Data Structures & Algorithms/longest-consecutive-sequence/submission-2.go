func longestConsecutive(nums []int) int {
	numsDir := make(map[int]bool)

	for _, num := range nums { // build the hashmap
		numsDir[num] = false
	}

	longest := min(1, len(nums));
	for _, num := range nums {
		if numsDir[num] {
			continue
		}

		_, ok := numsDir[num - 1]
		if !ok {
			numsDir[num] = true
			currentLongest, seekNumber := 1, num + 1
			for currentLongest <= len(nums) {
				_, ok := numsDir[seekNumber]
				if ok {
					numsDir[seekNumber] = true
					currentLongest++
					seekNumber++
				} else {
					break
				}
			}
			longest = max(longest, currentLongest)
		}
	}

	return longest
}
