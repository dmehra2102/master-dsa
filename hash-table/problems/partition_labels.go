package hashtable

func PartitionLabels(s string) []int {
	// STEP-1 : Create a hash table
	hashTable := map[rune]int{}

	// STEP-2 : Loop through an entire string and store the last occurence of every letter inside hash table
	for index, value := range s {
		hashTable[value] = index
	}

	// STEP-3 : create result array and with the help of two pointer find out min length of partition
	result := []int{}
	start := 0
	end := 0
	for i, value := range s {
		index := hashTable[value]
		if index > end {
			end = index
		}
		if i == end {
			result = append(result, i-start+1)
			start = i + 1
		}
	}

	return result
}
