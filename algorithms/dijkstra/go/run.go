package main

import "fmt"

func main() {
	graph := map[int][]Edge{
		1: {{To: 2, Weight: 1}, {To: 3, Weight: 4}},
		2: {{To: 3, Weight: 2}, {To: 4, Weight: 5}},
		3: {{To: 4, Weight: 1}},
		4: {},
	}
	dist := Dijkstra(graph, 1)
	fmt.Println("Dijkstra(graph, 1):")
	for node, d := range dist {
		fmt.Printf("  node %d: %d\n", node, d)
	}
}
