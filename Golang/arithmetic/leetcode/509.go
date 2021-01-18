package main

import "fmt"

//带备忘录的递归     斐波那契
func fib(N int) int {
	if N < 1 {
		return 0
	}
	//备忘录
	memo := make([]int, N+1, N+1)
	//
	return helper(memo, N)
}

func helper(memo []int, N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	if memo[N] != 0 {
		return memo[N]
	}
	memo[N] = helper(memo, N-1) + helper(memo, N-2)
	return memo[N]

}

//DP table  迭代
func fib2(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	dp := make([]int, N+1, N+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]

}

func main() {
	fmt.Println(fib(20))
	fmt.Println(fib2(20))
}
