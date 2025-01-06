package array

import "sort"

// This is a good problem
func MergeIntervals(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if prev[1] >= current[0] {
			if current[1] > prev[1] {
				prev[1] = current[1]
			}
		} else {
			result = append(result, prev)
			prev = current
		}
	}

	result = append(result, prev)
	return result
}
