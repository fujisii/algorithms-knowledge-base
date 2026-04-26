package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type heapEntry struct {
	node int
	dist int
}

type minHeap []heapEntry

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any) { *h = append(*h, x.(heapEntry)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// グラフ定義に現れる全ノード（到達不可能なものを含む）を収集する。
// 戻り値の dist には未到達ノードも math.MaxInt で含まれる仕様。
func collectNodes(graph map[int][]Edge, start int) map[int]bool {
	nodes := map[int]bool{start: true}
	for node, edges := range graph {
		nodes[node] = true
		for _, e := range edges {
			nodes[e.To] = true
		}
	}
	return nodes
}

func Dijkstra(graph map[int][]Edge, start int) map[int]int {
	for _, edges := range graph {
		for _, e := range edges {
			if e.Weight < 0 {
				panic(fmt.Errorf("negative edge weight %d is not supported", e.Weight))
			}
		}
	}
	nodes := collectNodes(graph, start)
	dist := make(map[int]int, len(nodes))
	for node := range nodes {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &minHeap{{node: start, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		entry := heap.Pop(pq).(heapEntry)
		u, du := entry.node, entry.dist
		if du > dist[u] {
			continue
		}
		for _, e := range graph[u] {
			if du > math.MaxInt-e.Weight {
				continue
			}
			newDist := du + e.Weight
			if newDist < dist[e.To] {
				dist[e.To] = newDist
				heap.Push(pq, heapEntry{node: e.To, dist: newDist})
			}
		}
	}

	return dist
}
