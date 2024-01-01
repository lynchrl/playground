package dpfib

func fibSlow(n int) int {
	if n <= 2 {
		return 1
	}
	return fibSlow(n-1) + fibSlow(n-2)
}

func fibMemo(n int) int {
	memo := make(map[int]int)
	return fibMemoHelper(n, memo)
}

func fibMemoHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibMemoHelper(n-1, memo) + fibMemoHelper(n-2, memo)
	return memo[n]
}
