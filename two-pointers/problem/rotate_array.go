package twopointers

/*
The rotate function modifies the nums array in place because you are passing a slice to the function, not an array. In Go, slices are references to an underlying array, which means that when you modify the elements of a slice, you are directly modifying the elements of the original array.
*/

func RotateArray(nums []int, k int) {
	// STEP-1 : Reverse the whole slice
	reverse(nums, 0, len(nums)-1)

	// STEP-2 : Find the nuber of times we actually have to rotate
	mod := k % len(nums)

	// STEP-3 : With the help of 2 pointer reverse first k element and then last (n-k) element
	start := 0
	end := mod - 1
	reverse(nums, start, end)

	start = mod
	end = len(nums) - 1
	reverse(nums, start, end)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
