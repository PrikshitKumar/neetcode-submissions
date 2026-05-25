func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
	
		stack = append(stack, i)
	}

	return result
}
