package array

func RomanToInt(s string) int {
	// STEP-1: Create a hash table and store the symbol with value.
	hashTable := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500,
		'M': 1000,
	}

	result := 0
	length := len(s)

	for i := 0; i < length; i++ {
		// Check if this is not the last character and if the current value is less than the next value
		if i < length-1 && hashTable[rune(s[i])] < hashTable[rune(s[i+1])] {
			result -= hashTable[rune(s[i])]
		} else {
			result += hashTable[rune(s[i])]
		}
	}

	return result
}
