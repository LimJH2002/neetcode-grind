func lengthOfLongestSubstring(s string) int {
	exists := make(map[byte] bool)
	left, right := 0, 0
	longest := 0

	for right < len(s) {
		char := s[right]
		length := right - left + 1

		if exists[char] != true { // current char doesn't exists
			exists[char] = true
			right++
		} else {
			for exists[char] {
				exists[s[left]] = false
				left++
				length--
			}
		}

		longest = max(longest, length)
	}

	return longest
}
