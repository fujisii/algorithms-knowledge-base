export function fibonacci(n: number): number {
  if (n === 0) return 0
  if (n === 1) return 1

  const dp: number[] = new Array(n + 1)
  dp[0] = 0
  dp[1] = 1

  for (let i = 2; i <= n; i++) {
    dp[i] = dp[i - 1] + dp[i - 2]
  }

  return dp[n]
}
