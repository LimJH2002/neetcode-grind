func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1

	for left < right {
		if !isAlphanumeric(rune(s[left])) {
			left++
			continue
		}

		if !isAlphanumeric(rune(s[right])) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}
