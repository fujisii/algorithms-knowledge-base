package main

// 早期終了最適化（内側ループでスワップが0回なら break）は省略。
// 最悪ケースの O(n²) を常に実行する素朴な実装。
func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}
