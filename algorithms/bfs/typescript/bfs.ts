export function bfs(graph: Map<number, number[]>, start: number): number[] {
  const visited = new Set<number>()
  const queue: number[] = [start]
  const result: number[] = []

  visited.add(start)

  while (queue.length > 0) {
    const node = queue.shift()!
    result.push(node)

    const neighbors = graph.get(node) ?? []
    for (const neighbor of neighbors) {
      if (!visited.has(neighbor)) {
        visited.add(neighbor)
        queue.push(neighbor)
      }
    }
  }

  return result
}
