package hashtable

import "sort"

func MinSetSize(arr []int) int {
	hashTable := map[int]int{}

	for _, val := range arr {
		hashTable[val]++
	}

	// Storing the frequency inside a slice
	freqSlice := make([]int, 0, len(hashTable))
	for _, count := range hashTable {
		freqSlice = append(freqSlice, count)
	}

	// Sort frequencies in descending order
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i] > freqSlice[j]
	})

	halfSize := len(arr) / 2
	removedElementsCount := 0
	setSize := 0

	for _, count := range freqSlice {
		removedElementsCount += count
		setSize++
		if removedElementsCount >= halfSize {
			break
		}
	}
	return setSize
}
