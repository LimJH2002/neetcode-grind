func minWindow(s string, t string) string {
	essential := make(map[byte]int)
	for _, c := range t {
		essential[byte(c)]++
	}

	left := 0
	res, found := "", false
	required, current := len(essential), 0
	window := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if window[c] == essential[c] {
			current++
		}

		for required == current {
			if !found || len(res) > right-left+1 {
				found = true
				res = s[left : right+1]
			}

			lc := s[left]
			window[lc]--
			if window[lc] < essential[lc] {
				current--
			}
			left++
		}
	}

	return res
}