package main

import (
	"strings"
	"testing"
)

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
		{"正常系: n=50", 50, 12586269025},
		// F(93) 以降は int64 の符号付きオーバーフローが発生する
		{"境界: n=92 は int64 範囲内で正確", 92, 7540113804746346429},
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
		r := recover()
		if r == nil {
			t.Error("n < 0 でパニックが発生するべき")
			return
		}
		msg, ok := r.(string)
		if !ok {
			t.Errorf("パニック値が string ではない: %T %v", r, r)
			return
		}
		if !strings.Contains(msg, "n must be >= 0") {
			t.Errorf("予期しないパニックメッセージ: %v", msg)
		}
	}()
	Fibonacci(-1)
}
