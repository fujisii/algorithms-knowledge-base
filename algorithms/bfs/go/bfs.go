package main

func BFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	dequeueIndex := 0
	result := []int{}

	for dequeueIndex < len(queue) {
		node := queue[dequeueIndex]
		dequeueIndex++
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}
