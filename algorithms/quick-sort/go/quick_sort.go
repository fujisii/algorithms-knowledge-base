package main

// 関数型スタイル実装（可読性優先）。スライスを毎回生成するため
// 空間計算量は O(n log n)。インプレース版（O(log n)）とは異なる点に注意。
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return append([]int{}, arr...)
	}

	pivot := arr[len(arr)/2]
	var left, middle, right []int

	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		default:
			right = append(right, v)
		}
	}

	result := QuickSort(left)
	result = append(result, middle...)
	result = append(result, QuickSort(right)...)
	return result
}
