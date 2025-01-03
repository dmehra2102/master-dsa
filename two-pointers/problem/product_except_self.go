package twopointers

/*
- Array Size Must Be Constant: In Go, arrays are fixed-size and their size must be known at compile time. You cannot use a variable (like N) to define the size of an array.
- Using Slices Instead: If you need a dynamically sized collection that can change during runtime, you should use slices instead of arrays.
*/

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
