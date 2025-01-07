package binarysearch

// SOLUTION-1  => T.C : O(n)
func FindPeakElement1(nums []int) int {
	N := len(nums)

	if N == 1 {
		return 0
	}
	for i := 0; i < N; i++ {
		if i == 0 && nums[i] > nums[i+1] {
			return i
		} else if i == N-1 && nums[i] > nums[i-1] {
			return i
		} else if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
	}

	return -1
}

// SOLUTION-2 => T.C : O(logN)
func DindPeakElement2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
