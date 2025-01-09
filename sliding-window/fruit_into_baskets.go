package slidingwindow

func TotalFruit(fruits []int) int {
	hashTable := map[int]int{}
	j := 0
	result := 0

	for i, val := range fruits {
		hashTable[val]++

		for len(hashTable) > 2 {
			leftItem := fruits[j]
			hashTable[leftItem]--
			if hashTable[leftItem] == 0 {
				delete(hashTable, leftItem)
			}
			j++
		}

		result = max(result, i-j+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
