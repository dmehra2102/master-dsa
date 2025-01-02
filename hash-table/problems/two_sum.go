package hashtable

func TwoSum(nums []int, target int) []int {
	// STEP-1 : Create a HashTable
	hashTable := map[int]int{}

	// STEP-2 : Run a loop in an nums
	for index, v := range nums {
		diff := target - v
		if value, exists := hashTable[diff]; exists {
			return []int{value, index}
		}
		hashTable[v] = index
	}

	return []int{-1, -1}
}
