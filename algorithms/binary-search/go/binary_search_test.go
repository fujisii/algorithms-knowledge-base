package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"正常系: 中央", []int{1, 3, 5, 7, 9}, 5, 2},
		{"正常系: 左端", []int{1, 3, 5, 7, 9}, 1, 0},
		{"正常系: 右端", []int{1, 3, 5, 7, 9}, 9, 4},
		{"境界: 空スライス", []int{}, 1, -1},
		{"境界: 単一要素一致", []int{5}, 5, 0},
		{"境界: 単一要素不一致", []int{5}, 3, -1},
		{"異常系: 未検出", []int{1, 3, 5, 7, 9}, 4, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
