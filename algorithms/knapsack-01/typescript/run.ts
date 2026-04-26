import { knapsack } from './knapsack-01'

const items = [
  { weight: 2, value: 6 },
  { weight: 2, value: 10 },
  { weight: 3, value: 12 },
]
const result = knapsack(items, 5)
console.log(result)
