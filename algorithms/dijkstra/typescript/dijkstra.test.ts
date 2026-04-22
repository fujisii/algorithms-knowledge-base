import { describe, it, expect } from 'vitest'
import { dijkstra, type WeightedGraph } from './dijkstra'

describe('dijkstra', () => {
  it('正常系: 重み付きグラフの最短距離を求める', () => {
    const graph: WeightedGraph = new Map([
      [1, [{ to: 2, weight: 1 }, { to: 3, weight: 4 }]],
      [2, [{ to: 3, weight: 2 }, { to: 4, weight: 5 }]],
      [3, [{ to: 4, weight: 1 }]],
      [4, []],
    ])
    const dist = dijkstra(graph, 1)
    expect(dist.get(1)).toBe(0)
    expect(dist.get(2)).toBe(1)
    expect(dist.get(3)).toBe(3)
    expect(dist.get(4)).toBe(4)
  })

  it('正常系: 単一ノードは距離 0 を返す', () => {
    const graph: WeightedGraph = new Map([[1, []]])
    const dist = dijkstra(graph, 1)
    expect(dist.get(1)).toBe(0)
  })

  it('境界: 到達不可能ノードは Infinity を返す', () => {
    const graph: WeightedGraph = new Map([
      [1, [{ to: 2, weight: 1 }]],
      [2, []],
      [3, []],
    ])
    const dist = dijkstra(graph, 1)
    expect(dist.get(1)).toBe(0)
    expect(dist.get(2)).toBe(1)
    expect(dist.get(3)).toBe(Infinity)
  })

  it('境界: スタートノードが隣接リストに存在しない', () => {
    const graph: WeightedGraph = new Map()
    const dist = dijkstra(graph, 1)
    expect(dist.get(1)).toBe(0)
  })
})
