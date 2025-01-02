package hashtable

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	// STEP-1 : Sort the array and store the result in new variable
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	// STEP-2 : Create a hash table
	hashTable := map[int]int{}

	// STEP-3 : Loop through an sorted array and store the index value of it inside it
	for key, val := range sorted {
		if _, exists := hashTable[val]; !exists {
			hashTable[val] = key
		}
	}

	// STEP-4 : loop through a nums array and started putting count inside result array
	result := make([]int, len(nums))
	for index, val := range nums {
		count := hashTable[val]
		result[index] = count
	}

	return result
}
