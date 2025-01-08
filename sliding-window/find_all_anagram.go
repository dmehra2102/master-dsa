package slidingwindow

func FindAnagrams(s string, p string) []int {
	n := len(s)
	m := len(p)
	if n < m {
		return []int{}
	}

	hashTable := map[rune]int{}
	for index := range p {
		hashTable[rune(p[index])]++
	}

	result := []int{}
	windowFreq := map[rune]int{}

	for i := 0; i < m; i++ {
		windowFreq[rune(s[i])]++
	}

	if equalMaps(hashTable, windowFreq) {
		result = append(result, 0)
	}

	for i := m; i < n; i++ {
		windowFreq[rune(s[i])]++   // just adding the new element
		windowFreq[rune(s[i-m])]-- // just removing the leftmost element

		if windowFreq[rune(s[i-m])] == 0 {
			delete(windowFreq, rune(s[i-m]))
		}

		if equalMaps(hashTable, windowFreq) {
			result = append(result, i-m+1)
		}

	}
	return result
}

func equalMaps(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}
