func maxArea(heights []int) int {
	area := 0;
	left, right := 0, len(heights) - 1

	for left < right {
		currentArea := (right - left) * min(heights[left], heights[right])
		area = max(area, currentArea)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return area
}
