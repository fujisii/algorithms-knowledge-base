package main

func DFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{}
	result := []int{}

	// 深いグラフでは goroutine スタックが大きく拡張される（通常は自動拡張だが、非常に深い場合はスタックオーバーフローになりうる）
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
