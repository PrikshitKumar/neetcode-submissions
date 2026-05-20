func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	var maxArea int = 0

	left, right := 0, len(heights)-1
	for left < right {
		height := min(heights[left], heights[right])
		width := right - left
		area := height * width

		if maxArea < area {
			maxArea = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
