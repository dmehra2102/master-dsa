package binarysearch

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	start := 0
	end := len(nums) - 1
	tar_ind := -1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			tar_ind = mid
			break
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if tar_ind == -1 {
		return result
	}

	j := tar_ind
	// Find First occurance
	for start < j {
		if nums[j-1] == target {
			j--
		} else {
			break
		}
	}

	i := tar_ind
	// Find last occurance
	for i < end {
		if nums[i+1] == target {
			i++
		} else {
			break
		}
	}

	return []int{j, i}
}
