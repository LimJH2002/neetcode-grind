func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	sMap := make(map[byte] int)
	tMap := make(map[byte] int)

	// Populating
	for i := 0; i < len(s); i++ {
		sMap[s[i]] = sMap[s[i]] + 1
		tMap[t[i]] = tMap[t[i]] + 1
	}

	// Checking
	for char, sFreq := range sMap {
		tFreq := tMap[char]
		if (sFreq != tFreq) {
			return false
		}
	}

	return true
}
