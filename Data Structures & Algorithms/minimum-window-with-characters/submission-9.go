func minWindow(s string, t string) string {
	essential := make(map[byte]int)
	for _, c := range t {
		essential[byte(c)]++
	}

	res := ""
	required, current, left := len(essential), 0, 0
	window := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if window[c] == essential[c] {
			current++
		}

		for required == current {
			if res == "" || len(res) > right-left+1 {
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