package recursion

func Fib(n int) int {
	memo := make(map[int]int)

	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}
