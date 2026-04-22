import { bfs } from './bfs'

const graph = new Map([
  [1, [2, 3]],
  [2, [4]],
  [3, [4]],
  [4, []],
])
const result = bfs(graph, 1)
console.log(`bfs(graph, 1) = [${result}]`)
