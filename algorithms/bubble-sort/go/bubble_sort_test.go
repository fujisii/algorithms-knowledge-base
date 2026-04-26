package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"正常系: 整数配列", []int{5, 3, 1, 4, 2}, []int{1, 2, 3, 4, 5}},
		{"境界: 空スライス", []int{}, []int{}},
		{"境界: 単一要素", []int{42}, []int{42}},
		{"境界: 既ソート済み", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"境界: 逆順", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"境界: 重複あり", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"境界: 負数を含む", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BubbleSort(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
