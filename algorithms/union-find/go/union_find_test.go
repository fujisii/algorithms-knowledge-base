package main

import "testing"

func TestUnionFind(t *testing.T) {
	t.Run("初期状態: 異なる要素は未接続", func(t *testing.T) {
		uf := NewUnionFind(5)
		if uf.Connected(0, 1) {
			t.Error("0と1は未接続であるべき")
		}
		if uf.Connected(2, 4) {
			t.Error("2と4は未接続であるべき")
		}
	})

	t.Run("union後: 結合された要素はconnected", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		if !uf.Connected(0, 1) {
			t.Error("0と1は接続されるべき")
		}
	})

	t.Run("推移的接続: 中間ノードを経由して接続される", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		uf.Union(1, 2)
		if !uf.Connected(0, 2) {
			t.Error("0と2は接続されるべき")
		}
	})

	t.Run("別グループ: union されていない要素は未接続", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		uf.Union(3, 4)
		if uf.Connected(0, 3) {
			t.Error("0と3は未接続であるべき")
		}
	})

	t.Run("全接続: 連鎖union後すべての要素が接続される", func(t *testing.T) {
		uf := NewUnionFind(4)
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)
		pairs := []struct{ x, y int }{{0, 3}, {1, 3}, {0, 2}}
		for _, p := range pairs {
			if !uf.Connected(p.x, p.y) {
				t.Errorf("%dと%dは接続されるべき", p.x, p.y)
			}
		}
	})

	t.Run("自己参照: 同じ要素同士は常にconnected", func(t *testing.T) {
		uf := NewUnionFind(3)
		if !uf.Connected(0, 0) {
			t.Error("0は自分自身と接続されるべき")
		}
	})
}

func TestUnionFindPanic(t *testing.T) {
	uf := NewUnionFind(3)
	panics := []struct {
		name string
		idx  int
	}{
		{"上限超過", 5},
		{"負のインデックス", -1},
	}
	for _, tc := range panics {
		tc := tc
		t.Run(tc.name+"でパニック", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Find(%d) でパニックが発生するべき", tc.idx)
				}
			}()
			uf.Find(tc.idx)
		})
	}
}
