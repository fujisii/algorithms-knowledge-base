import { dfs } from './dfs'

const graph = new Map([
  [1, [2, 3]],
  [2, [4]],
  [3, [4]],
  [4, []],
])
const result = dfs(graph, 1)
console.log(`dfs(graph, 1) = [${result}]`)
