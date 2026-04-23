package main

import (
	"math"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]Edge
		start int
		want  map[int]int
	}{
		{
			name: "正常系: 重み付きグラフの最短距離を求める",
			graph: map[int][]Edge{
				1: {{To: 2, Weight: 1}, {To: 3, Weight: 4}},
				2: {{To: 3, Weight: 2}, {To: 4, Weight: 5}},
				3: {{To: 4, Weight: 1}},
				4: {},
			},
			start: 1,
			want:  map[int]int{1: 0, 2: 1, 3: 3, 4: 4},
		},
		{
			name:  "正常系: 単一ノードは距離 0 を返す",
			graph: map[int][]Edge{1: {}},
			start: 1,
			want:  map[int]int{1: 0},
		},
		{
			name: "境界: 到達不可能ノードは MaxInt を返す",
			graph: map[int][]Edge{
				1: {{To: 2, Weight: 1}},
				2: {},
				3: {},
			},
			start: 1,
			want:  map[int]int{1: 0, 2: 1, 3: math.MaxInt},
		},
		{
			name:  "境界: スタートノードが隣接リストに存在しない",
			graph: map[int][]Edge{},
			start: 1,
			want:  map[int]int{1: 0},
		},
		{
			name: "正常系: 等コストの並列経路で最短距離を正しく選ぶ",
			// 1→2→4 (1+3=4) と 1→3→4 (2+2=4) が同コスト
			graph: map[int][]Edge{
				1: {{To: 2, Weight: 1}, {To: 3, Weight: 2}},
				2: {{To: 4, Weight: 3}},
				3: {{To: 4, Weight: 2}},
				4: {},
			},
			start: 1,
			want:  map[int]int{1: 0, 2: 1, 3: 2, 4: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dijkstra(tt.graph, tt.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestDijkstraNegativeWeightPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("負の辺の重みでパニックが発生するべき")
		}
	}()
	graph := map[int][]Edge{
		1: {{To: 2, Weight: -1}},
	}
	Dijkstra(graph, 1)
}
