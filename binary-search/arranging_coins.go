package binarysearch

// Very good problem
func ArrangeCoins(n int) int {

	start := 1
	end := n
	for start <= end {
		mid := start + (end-start)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		} else if sum > n {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return end
}
