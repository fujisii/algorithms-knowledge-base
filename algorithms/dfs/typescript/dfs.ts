export function dfs(graph: Map<number, number[]>, start: number): number[] {
  const visited = new Set<number>()
  const result: number[] = []

  function visit(node: number): void {
    if (visited.has(node)) return
    visited.add(node)
    result.push(node)
    for (const neighbor of graph.get(node) ?? []) {
      visit(neighbor)
    }
  }

  visit(start)
  return result
}
