func evalRPN(tokens []string) int {
	stack := []int{}

	for _, val := range tokens {
		if val != "+" && val != "-" && val != "*" && val != "/" {
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			
			var result int

			switch val {
				case "+":
					result = num1 + num2
				case "-": 
					result = num1 - num2
				case "*": 
					result = num1 * num2
				case "/": 
					result = num1 / num2
			}

			stack = append(stack, result)
		}
	}

	return stack[0]
}
