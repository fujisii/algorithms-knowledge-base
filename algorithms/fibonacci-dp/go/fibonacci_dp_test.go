package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"境界: n=0", 0, 0},
		{"境界: n=1", 1, 1},
		{"正常系: n=2", 2, 1},
		{"正常系: n=10", 10, 55},
		{"正常系: n=20", 20, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fibonacci(tt.n)
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("n < 0 でパニックが発生するべき")
		}
	}()
	Fibonacci(-1)
}
