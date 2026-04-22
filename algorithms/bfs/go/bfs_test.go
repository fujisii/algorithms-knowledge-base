package main

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]int
		start int
		want  []int
	}{
		{
			name:  "正常系: 連結グラフを幅優先順に訪問する",
			graph: map[int][]int{1: {2, 3}, 2: {4}, 3: {4}, 4: {}},
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "正常系: 単一ノードは自身のみ返す",
			graph: map[int][]int{1: {}},
			start: 1,
			want:  []int{1},
		},
		{
			name:  "境界: 到達不可能ノードは含まない",
			graph: map[int][]int{1: {2}, 2: {}, 3: {}},
			start: 1,
			want:  []int{1, 2},
		},
		{
			name:  "境界: 閉路ありグラフで無限ループしない",
			graph: map[int][]int{1: {2}, 2: {3}, 3: {1}},
			start: 1,
			want:  []int{1, 2, 3},
		},
		{
			name:  "境界: スタートノードが隣接リストに存在しない",
			graph: map[int][]int{},
			start: 1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BFS(tt.graph, tt.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v; want %v", got, tt.want)
			}
		})
	}
}
