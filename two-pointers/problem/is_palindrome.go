package twopointers

import "unicode"

func IsPalindrome(s string) bool {
	validStr := modifyString(s)

	start, end := 0, len(validStr)-1

	for start < end {
		if validStr[start] != validStr[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func modifyString(s string) string {
	result := []rune{}
	for _, val := range s {
		if isTrue := unicode.IsLetter(val); isTrue {
			result = append(result, unicode.ToLower(val))
		} else if isTrue := unicode.IsDigit(val); isTrue {
			result = append(result, val)
		}
	}

	return string(result)
}
