func lengthOfLongestSubstring(s string) int {
	exists := make(map[byte] bool)
	left, longest := 0, 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		
		for exists[char] {
			exists[s[left]] = false
			left++
		}
		exists[char] = true
		longest = max(longest, right - left + 1)
	}

	return longest
}
