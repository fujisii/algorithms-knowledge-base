export type Item = { weight: number; value: number }

export function knapsack(items: Item[], capacity: number): number {
  if (capacity < 0) throw new RangeError(`capacity must be >= 0, got ${capacity}`)
  for (const item of items) {
    if (item.weight < 0) throw new RangeError(`item weight must be >= 0, got ${item.weight}`)
  }
  const n = items.length
  const dp: number[][] = Array.from({ length: n + 1 }, () => new Array(capacity + 1).fill(0))

  for (let i = 1; i <= n; i++) {
    const { weight, value } = items[i - 1]
    for (let w = 0; w <= capacity; w++) {
      dp[i][w] = dp[i - 1][w]
      if (w >= weight) {
        dp[i][w] = Math.max(dp[i][w], dp[i - 1][w - weight] + value)
      }
    }
  }

  return dp[n][capacity]
}
