package array

func IsValidParentheses(s string) bool {
	// Step 1: Create a stack for storing open brackets
	stack := []rune{}

	// Step 2: Start running a loop
	for _, val := range s {
		switch val {
		case '(', '{', '[':
			stack = append(stack, val)

		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]

		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
