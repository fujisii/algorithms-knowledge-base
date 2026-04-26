package main

import "fmt"

func main() {
	arr := []int{5, 3, 1, 4, 2}
	result := BubbleSort(arr)
	fmt.Printf("BubbleSort(%v) = %v\n", arr, result)
}
