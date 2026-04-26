// 関数型スタイル実装（可読性優先）。filter で分割配列を毎回生成するため
// 空間計算量は O(n log n)。インプレース版（O(log n)）とは異なる点に注意。
export function quickSort(arr: number[]): number[] {
  if (arr.length <= 1) return [...arr]

  const pivot = arr[Math.floor(arr.length / 2)]
  const left = arr.filter((x) => x < pivot)
  const middle = arr.filter((x) => x === pivot)
  const right = arr.filter((x) => x > pivot)

  return [...quickSort(left), ...middle, ...quickSort(right)]
}
