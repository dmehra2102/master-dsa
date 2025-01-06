package hashtable

import "sort"

func GroupAnagrams(strs []string) [][]string {
	hashTable := map[string][]string{}

	for _, val := range strs {
		runes := []rune(val)
		// Sort the runes
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		// Convert sorted runes back to a string
		sorted_str := string(runes)

		if _, exists := hashTable[sorted_str]; !exists {
			hashTable[sorted_str] = []string{}
		}
		hashTable[sorted_str] = append(hashTable[sorted_str], val)
	}

	result := make([][]string, 0, len(hashTable))
	for _, group := range hashTable {
		result = append(result, group)
	}

	return result
}
