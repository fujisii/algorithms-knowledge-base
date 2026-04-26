package main

import "fmt"

func main() {
	graph := map[int][]int{1: {2, 3}, 2: {4}, 3: {4}, 4: {}}
	result := BFS(graph, 1)
	fmt.Printf("BFS(graph, 1) = %v\n", result)
}
