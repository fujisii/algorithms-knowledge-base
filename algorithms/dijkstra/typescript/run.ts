import { dijkstra, type WeightedGraph } from './dijkstra'

const graph: WeightedGraph = new Map([
  [1, [{ to: 2, weight: 1 }, { to: 3, weight: 4 }]],
  [2, [{ to: 3, weight: 2 }, { to: 4, weight: 5 }]],
  [3, [{ to: 4, weight: 1 }]],
  [4, []],
])
const dist = dijkstra(graph, 1)
console.log('dijkstra(graph, 1):')
for (const [node, d] of dist) {
  console.log(`  node ${node}: ${d}`)
}
