package array

func SingleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res = res ^ val // using XOR(^) operator
	}

	return res
}

/*
The XOR operation has two important properties:
 -> Self-Canceling: For any integer a, a⊕a=0. This means that if you XOR a number with itself, the result is zero.
 -> Identity with Zero: For any integer a, a⊕0=a. This means that XORing any number with zero returns that number.
*/
