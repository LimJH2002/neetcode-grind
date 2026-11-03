func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	left, maxFreq, result := 0, 0, 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		freqMap[char]++
		maxFreq = max(maxFreq, freqMap[char])

		if (right - left + 1) - maxFreq > k {
			freqMap[s[left]]--
			left++
		}

		currentMax := 0
		for _, freq := range freqMap {
			currentMax = max(currentMax, freq)
		}

		result = max(result, right - left + 1)
	}

	return result
}
