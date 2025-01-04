package hashtable

func Intersect(nums1 []int, nums2 []int) []int {
	hashTable := make(map[int]int)
	result := []int{}

	// Populate the hash table with counts from the smaller array
	if len(nums1) < len(nums2) {
		for _, val := range nums1 {
			hashTable[val]++
		}
		for _, val := range nums2 {
			if count, exists := hashTable[val]; exists && count > 0 {
				result = append(result, val)
				hashTable[val]-- // Decrease count after adding to result
			}
		}
	} else {
		for _, val := range nums2 {
			hashTable[val]++
		}
		for _, val := range nums1 {
			if count, exists := hashTable[val]; exists && count > 0 {
				result = append(result, val)
				hashTable[val]-- // Decrease count after adding to result
			}
		}
	}

	return result
}
