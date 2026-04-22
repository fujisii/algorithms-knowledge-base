package main

func DFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{}
	result := []int{}

	var visit func(node int)
	visit = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		result = append(result, node)
		for _, neighbor := range graph[node] {
			visit(neighbor)
		}
	}

	visit(start)
	return result
}
