package main

import "fmt"

func main() {
	for _, n := range []int{0, 1, 2, 5, 10} {
		fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n))
	}
}
