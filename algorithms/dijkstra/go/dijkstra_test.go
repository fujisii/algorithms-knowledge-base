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
