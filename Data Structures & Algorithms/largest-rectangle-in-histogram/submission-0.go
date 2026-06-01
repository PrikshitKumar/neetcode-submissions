func largestRectangleArea(heights []int) int {
	stack := [][]int{}
	maxArea := 0

	for i, h := range heights {
		start := i

		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			index := stack[len(stack)-1][0]
			height := stack[len(stack)-1][1]

			stack = stack[:len(stack)-1]

			area := height * (i - index)

			if area > maxArea {
				maxArea = area
			}

			start = index
		}

		stack = append(stack, []int{start, h})
	}

	for _, item := range stack {
		index := item[0]
		height := item[1]

		area := height * (len(heights) - index)

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
