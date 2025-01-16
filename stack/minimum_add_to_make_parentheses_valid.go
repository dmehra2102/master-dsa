package stack

// Solution 1
func MinAddToMakeValid1(s string) int {
	stack := []rune{}
	count := 0
	if len(s) == 0 {
		return 0
	}

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) == 0 {
				count++
			} else {
				stack = stack[1:]
			}
		}
	}

	return count + len(stack)
}

// Solution 2
func MinAddToMakeValid2(s string) int {
	openCount := 0
	closeCount := 0

	for _, char := range s {
		if char == '(' {
			openCount++
		} else if char == ')' {
			if openCount > 0 {
				openCount-- // Match with an opening parenthesis
			} else {
				closeCount++ // Unmatched closing parenthesis
			}
		}
	}

	// Total additions needed is unmatched opens + unmatched closes
	return openCount + closeCount
}
