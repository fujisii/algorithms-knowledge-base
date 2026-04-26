package main

import (
	"strings"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		items    []Item
		capacity int
		want     int
	}{
		{
			"正常系: 最適な組み合わせ",
			[]Item{{Weight: 2, Value: 6}, {Weight: 2, Value: 10}, {Weight: 3, Value: 12}},
			5,
			22,
		},
		{
			"正常系: 全アイテムが容量内",
			[]Item{{Weight: 1, Value: 1}, {Weight: 2, Value: 2}, {Weight: 3, Value: 3}},
			10,
			6,
		},
		{
			"境界: 空アイテムリスト",
			[]Item{},
			10,
			0,
		},
		{
			"境界: capacity=0",
			[]Item{{Weight: 1, Value: 10}},
			0,
			0,
		},
		{
			"境界: 全アイテムが容量超過",
			[]Item{{Weight: 5, Value: 10}, {Weight: 6, Value: 20}},
			4,
			0,
		},
		{
			"境界: 単一アイテムが丁度容量と等しい",
			[]Item{{Weight: 3, Value: 15}},
			3,
			15,
		},
		{
			"境界: weight=0 のアイテムは容量 0 でも選択可能",
			[]Item{{Weight: 0, Value: 5}},
			0,
			5,
		},
		{
			// 容量6に対して重さ3・価値10のアイテム1種: 2個取れるが0-1なので10
			"境界: 同一アイテムは1個のみ選択可能（0-1制約）",
			[]Item{{Weight: 3, Value: 10}},
			6,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Knapsack(tt.items, tt.capacity)
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestKnapsackPanic(t *testing.T) {
	tests := []struct {
		name    string
		items   []Item
		cap     int
		wantMsg string
	}{
		{"capacity < 0", []Item{}, -1, "capacity must be >= 0"},
		{"weight < 0", []Item{{Weight: -1, Value: 5}}, 10, "item weight must be >= 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					t.Error("パニックが発生するべき")
					return
				}
				err, ok := r.(error)
				if !ok {
					t.Errorf("パニック値が error ではない: %T %v", r, r)
					return
				}
				if !strings.Contains(err.Error(), tt.wantMsg) {
					t.Errorf("予期しないパニックメッセージ: %v", err)
				}
			}()
			Knapsack(tt.items, tt.cap)
		})
	}
}
