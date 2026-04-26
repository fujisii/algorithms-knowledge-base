export function bfs(graph: ReadonlyMap<number, readonly number[]>, start: number): number[] {
  const visited = new Set<number>()
  const queue: number[] = [start]
  let dequeueIndex = 0
  const result: number[] = []

  visited.add(start)

  while (dequeueIndex < queue.length) {
    const node = queue[dequeueIndex++]
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
