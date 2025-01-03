package twopointers

func SortArrayByParityII(nums []int) []int {
	// STEP-1 : Create 2 pointers
	i, j := 0, 1
	N := len(nums)

	// STEP-2 : Start running a two pointer
	for i < N && j < N {
		isEven := nums[i]%2 == 0
		if isEven {
			i += 2
		} else if !isEven && nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
			j += 2
		} else {
			j += 2
		}
	}
	return nums
}
