// 早期終了最適化（内側ループでスワップが0回なら break）は省略。
// 最悪ケースの O(n²) を常に実行する素朴な実装。
export function bubbleSort(arr: number[]): number[] {
  const result = [...arr]
  const n = result.length

  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - 1 - i; j++) {
      if (result[j] > result[j + 1]) {
        ;[result[j], result[j + 1]] = [result[j + 1], result[j]]
      }
    }
  }

  return result
}
