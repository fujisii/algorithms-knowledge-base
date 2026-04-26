package main

// graph は読み取り専用で使用する（Go に ReadonlyMap 相当はないため、変更しない前提）。
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
