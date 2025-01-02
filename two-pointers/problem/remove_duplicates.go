package twopointers

func RemoveDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}

	return i + 1
}
