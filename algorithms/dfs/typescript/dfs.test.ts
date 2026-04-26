import { describe, it, expect } from 'vitest'
import { dfs } from './dfs'

describe('dfs', () => {
  it('正常系: 連結グラフを深さ優先順に訪問する', () => {
    const graph = new Map([
      [1, [2, 3]],
      [2, [4]],
      [3, [4]],
      [4, []],
    ])
    expect(dfs(graph, 1)).toEqual([1, 2, 4, 3])
  })

  it('正常系: 単一ノードは自身のみ返す', () => {
    const graph = new Map([[1, []]])
    expect(dfs(graph, 1)).toEqual([1])
  })

  it('境界: 到達不可能ノードは含まない', () => {
    const graph = new Map([
      [1, [2]],
      [2, []],
      [3, []],
    ])
    expect(dfs(graph, 1)).toEqual([1, 2])
  })

  it('境界: 閉路ありグラフで無限ループしない', () => {
    const graph = new Map([
      [1, [2]],
      [2, [3]],
      [3, [1]],
    ])
    expect(dfs(graph, 1)).toEqual([1, 2, 3])
  })

  it('境界: スタートノードが隣接リストに存在しない', () => {
    const graph = new Map<number, number[]>()
    expect(dfs(graph, 1)).toEqual([1])
  })
})
