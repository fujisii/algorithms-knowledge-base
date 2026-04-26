export function binarySearch(arr: readonly number[], target: number): number {
  let left = 0
  let right = arr.length - 1

  while (left <= right) {
    // number は IEEE 754 倍精度浮動小数点のため整数オーバーフローは発生しない（Go 実装では left+(right-left)/2 でオーバーフロー回避）
    const mid = Math.floor((left + right) / 2)
    if (arr[mid] === target) return mid
    if (arr[mid] < target) left = mid + 1
    else right = mid - 1
  }

  return -1
}
