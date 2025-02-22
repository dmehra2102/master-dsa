package twopointers

func TwoSumII(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum < target {
			start++
		} else {
			end--
		}
	}

	return []int{start + 1, end + 1}
}
