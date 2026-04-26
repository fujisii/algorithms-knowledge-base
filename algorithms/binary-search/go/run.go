package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 5
	result := BinarySearch(arr, target)
	fmt.Printf("BinarySearch(%v, %d) = %d\n", arr, target, result)
}
