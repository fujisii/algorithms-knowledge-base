package main

import (
	"container/heap"
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
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(heapEntry)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

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
