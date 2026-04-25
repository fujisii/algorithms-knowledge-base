package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"正常系: 一般的な配列", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"正常系: 既ソート済み", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"正常系: 逆順", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"境界: 空スライス", []int{}, []int{}},
		{"境界: 単一要素", []int{42}, []int{42}},
		{"境界: 重複要素あり", []int{3, 3, 1, 1, 2, 2}, []int{1, 1, 2, 2, 3, 3}},
		{"境界: 2要素", []int{2, 1}, []int{1, 2}},
		{"境界: 負数を含む", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
