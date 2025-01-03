package twopointers

func ProductExceptSelf(nums []int) []int {
	N := len(nums)
	result := make([]int, N)

	// STEP-1 : start a Loop from left to right and store the product of left-side element
	product := 1
	for index := 0; index < N; index++ {
		result[index] = product
		product *= nums[index]
	}

	// STEP-2 : Start a Loop from right to left and store the prodcut of right-side element
	product = 1
	for index := N - 1; index >= 0; index-- {
		result[index] = result[index] * product
		product *= nums[index]
	}

	return result
}
