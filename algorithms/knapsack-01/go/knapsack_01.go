package main

import "fmt"

type Item struct {
	Weight int
	Value  int
}

func Knapsack(items []Item, capacity int) int {
	if capacity < 0 {
		panic(fmt.Sprintf("capacity must be >= 0, got %d", capacity))
	}
	for _, item := range items {
		if item.Weight < 0 {
			panic(fmt.Sprintf("item weight must be >= 0, got %d", item.Weight))
		}
	}
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		item := items[i-1]
		for w := 0; w <= capacity; w++ {
			dp[i][w] = dp[i-1][w]
			if w >= item.Weight {
				candidate := dp[i-1][w-item.Weight] + item.Value
				if candidate > dp[i][w] {
					dp[i][w] = candidate
				}
			}
		}
	}

	return dp[n][capacity]
}
