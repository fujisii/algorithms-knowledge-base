package main

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
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
