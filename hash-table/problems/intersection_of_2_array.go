package hashtable

func Intersection(nums1 []int, nums2 []int) []int {
	// STEP-1 : Create a hashTable and store all the elements of sorter array
	hashTable := map[int]struct{}{}
	minArr, maxArr := minArray(nums1, nums2)
	for _, val := range minArr {
		hashTable[val] = struct{}{}
	}

	// STEP-2 Loop through the larger array and store the result
	result := make([]int, 0, len(minArr))
	for _, val := range maxArr {
		if _, exists := hashTable[val]; exists {
			result = append(result, val)
			delete(hashTable, val)
		}
	}

	return result
}

func minArray(nums1, nums2 []int) ([]int, []int) {
	if len(nums1) < len(nums2) {
		return nums1, nums2
	}
	return nums2, nums1
}
