package twopointers

// Very important
func BackspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		i = nextValidCharIndex(s, i)
		j = nextValidCharIndex(t, j)

		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}

		if (i >= 0) != (j >= 0) {
			return false
		}

		i--
		j--
	}
	return true
}

func nextValidCharIndex(s string, index int) int {
	backspaceCount := 0
	for index >= 0 {
		if s[index] == '#' {
			backspaceCount++
		} else {
			if backspaceCount > 0 {
				backspaceCount--
			} else {
				break
			}
		}
		index--
	}
	return index
}
