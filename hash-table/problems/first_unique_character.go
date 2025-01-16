package hashtable

func FirstUniqChar(s string) int {
	hashMap := make(map[rune]int)

	for i := range s {
		hashMap[rune(s[i])]++
	}

	for i := range s {
		val := hashMap[rune(s[i])]
		if val == 1 {
			return i
		}
	}

	return -1
}
