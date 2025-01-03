package twopointers

func SortColors(nums []int) {
	// STEP-1 : Create three pointers
	i, j, k := 0, 0, len(nums)-1

	// STEP-2 : Run a loop and use Dutch National Flag Algorithm
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 1 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}
