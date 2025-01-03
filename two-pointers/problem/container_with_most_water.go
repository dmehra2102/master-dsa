package twopointers

func MaxArea(height []int) int {
	// STEP-1 create 2 pointers and max variable
	start, end := 0, len(height)-1
	max := 0

	// STEP-2 : with the help two pointer you can easily solve this problem
	for start < end {
		minH := minHeight(height[start], height[end])
		max = maxHeight(max, minH*(end-start))

		if minH == height[start] {
			start++
		} else {
			end--
		}
	}

	return max
}

func maxHeight(height1 int, height2 int) int {
	if height1 > height2 {
		return height1
	}
	return height2
}

func minHeight(height1 int, height2 int) int {
	if height1 > height2 {
		return height2
	}
	return height1
}
