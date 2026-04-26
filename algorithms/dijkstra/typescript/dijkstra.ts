export type Edge = { to: number; weight: number }
export type WeightedGraph = Map<number, Edge[]>

type HeapEntry = { node: number; dist: number }

// グラフ定義に現れる全ノード（到達不可能なものを含む）を収集する。
// 戻り値の dist には未到達ノードも Infinity で含まれる仕様。
function collectNodes(graph: WeightedGraph, start: number): Set<number> {
  const nodes = new Set<number>([start])
  for (const [node, edges] of graph) {
    nodes.add(node)
    for (const { to } of edges) nodes.add(to)
  }
  return nodes
}

class MinHeap {
  private data: HeapEntry[] = []

  get size() {
    return this.data.length
  }

  push(entry: HeapEntry): void {
    this.data.push(entry)
    this.bubbleUp(this.data.length - 1)
  }

  pop(): HeapEntry | undefined {
    if (this.data.length === 0) return undefined
    const top = this.data[0]
    const last = this.data.pop()!
    if (this.data.length > 0) {
      this.data[0] = last
      this.sinkDown(0)
    }
    return top
  }

  private bubbleUp(i: number): void {
    while (i > 0) {
      const parent = (i - 1) >> 1
      if (this.data[parent].dist <= this.data[i].dist) break
      ;[this.data[parent], this.data[i]] = [this.data[i], this.data[parent]]
      i = parent
    }
  }

  private sinkDown(i: number): void {
    const n = this.data.length
    while (true) {
      let smallest = i
      const left = 2 * i + 1
      const right = 2 * i + 2
      if (left < n && this.data[left].dist < this.data[smallest].dist) smallest = left
      if (right < n && this.data[right].dist < this.data[smallest].dist) smallest = right
      if (smallest === i) break
      ;[this.data[smallest], this.data[i]] = [this.data[i], this.data[smallest]]
      i = smallest
    }
  }
}

export function dijkstra(graph: WeightedGraph, start: number): Map<number, number> {
  for (const edges of graph.values()) {
    for (const { weight } of edges) {
      if (weight < 0) throw new RangeError(`negative edge weight ${weight} is not supported`)
    }
  }
  const nodes = collectNodes(graph, start)
  const dist = new Map<number, number>()
  for (const node of nodes) dist.set(node, Infinity)
  dist.set(start, 0)

  const heap = new MinHeap()
  heap.push({ node: start, dist: 0 })

  while (heap.size > 0) {
    const { node: u, dist: du } = heap.pop()!
    if (du > dist.get(u)!) continue

    for (const { to, weight } of graph.get(u) ?? []) {
      const newDist = du + weight
      const toDist = dist.get(to)!
      if (newDist < toDist) {
        dist.set(to, newDist)
        heap.push({ node: to, dist: newDist })
      }
    }
  }

  return dist
}
