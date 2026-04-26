package main

import "fmt"

// n >= 93 では int（64bit 環境では int64 相当）の符号付きオーバーフローが発生し、誤った値を返す。
// 大きな n が必要な場合は math/big パッケージの使用を検討すること。
func Fibonacci(n int) int {
	if n < 0 {
		panic(fmt.Errorf("n must be >= 0, got %d", n))
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
