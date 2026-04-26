package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		validate func(t *testing.T, got int)
	}{
		{
			name:   "正常系: 中央",
			arr:    []int{1, 3, 5, 7, 9},
			target: 5,
			validate: func(t *testing.T, got int) {
				if got != 2 {
					t.Errorf("got %v; want 2", got)
				}
			},
		},
		{
			name:   "正常系: 左端",
			arr:    []int{1, 3, 5, 7, 9},
			target: 1,
			validate: func(t *testing.T, got int) {
				if got != 0 {
					t.Errorf("got %v; want 0", got)
				}
			},
		},
		{
			name:   "正常系: 右端",
			arr:    []int{1, 3, 5, 7, 9},
			target: 9,
			validate: func(t *testing.T, got int) {
				if got != 4 {
					t.Errorf("got %v; want 4", got)
				}
			},
		},
		{
			name:   "正常系: 重複要素あり（有効なインデックスのいずれかを返す）",
			arr:    []int{1, 3, 3, 3, 5},
			target: 3,
			validate: func(t *testing.T, got int) {
				if got < 1 || got > 3 {
					t.Errorf("got %v; want index in [1, 3]", got)
				}
			},
		},
		{
			name:   "境界: 空スライス",
			arr:    []int{},
			target: 1,
			validate: func(t *testing.T, got int) {
				if got != -1 {
					t.Errorf("got %v; want -1", got)
				}
			},
		},
		{
			name:   "境界: 単一要素一致",
			arr:    []int{5},
			target: 5,
			validate: func(t *testing.T, got int) {
				if got != 0 {
					t.Errorf("got %v; want 0", got)
				}
			},
		},
		{
			name:   "境界: 単一要素不一致",
			arr:    []int{5},
			target: 3,
			validate: func(t *testing.T, got int) {
				if got != -1 {
					t.Errorf("got %v; want -1", got)
				}
			},
		},
		{
			name:   "異常系: 未検出",
			arr:    []int{1, 3, 5, 7, 9},
			target: 4,
			validate: func(t *testing.T, got int) {
				if got != -1 {
					t.Errorf("got %v; want -1", got)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target)
			tt.validate(t, got)
		})
	}
}
