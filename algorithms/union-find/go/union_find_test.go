package main

import "testing"

func TestUnionFind_InitiallyDisconnected(t *testing.T) {
	uf := NewUnionFind(5)
	if uf.Connected(0, 1) {
		t.Error("初期状態: 0と1は未接続であるべき")
	}
	if uf.Connected(2, 4) {
		t.Error("初期状態: 2と4は未接続であるべき")
	}
}

func TestUnionFind_ConnectedAfterUnion(t *testing.T) {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	if !uf.Connected(0, 1) {
		t.Error("union後: 0と1は接続されるべき")
	}
}

func TestUnionFind_TransitiveConnection(t *testing.T) {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	if !uf.Connected(0, 2) {
		t.Error("推移的接続: 0と2は接続されるべき")
	}
}

func TestUnionFind_SeparateGroups(t *testing.T) {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(3, 4)
	if uf.Connected(0, 3) {
		t.Error("別グループ: 0と3は未接続であるべき")
	}
}

func TestUnionFind_AllConnected(t *testing.T) {
	uf := NewUnionFind(4)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	tests := []struct{ x, y int }{{0, 3}, {1, 3}, {0, 2}}
	for _, tt := range tests {
		if !uf.Connected(tt.x, tt.y) {
			t.Errorf("全接続: %dと%dは接続されるべき", tt.x, tt.y)
		}
	}
}

func TestUnionFind_SelfConnected(t *testing.T) {
	uf := NewUnionFind(3)
	if !uf.Connected(0, 0) {
		t.Error("自己参照: 0は自分自身と接続されるべき")
	}
}
