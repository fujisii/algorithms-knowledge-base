package main

import "fmt"

func main() {
	items := []Item{
		{Weight: 2, Value: 6},
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 12},
	}
	result := Knapsack(items, 5)
	fmt.Println(result)
}
