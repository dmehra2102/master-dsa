package array

func SingleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res = res ^ val // using XOR(^) operator
	}

	return res
}
