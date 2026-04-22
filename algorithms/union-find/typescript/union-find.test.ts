import { describe, it, expect } from 'vitest'
import { UnionFind } from './union-find'

describe('UnionFind', () => {
  it('初期状態: 全要素が独立している', () => {
    const uf = new UnionFind(5)
    expect(uf.connected(0, 1)).toBe(false)
    expect(uf.connected(2, 4)).toBe(false)
  })
  it('union後: 結合された要素はconnectedになる', () => {
    const uf = new UnionFind(5)
    uf.union(0, 1)
    expect(uf.connected(0, 1)).toBe(true)
  })
  it('union後: 推移的な接続が成立する', () => {
    const uf = new UnionFind(5)
    uf.union(0, 1)
    uf.union(1, 2)
    expect(uf.connected(0, 2)).toBe(true)
  })
  it('union後: 異なるグループは接続されない', () => {
    const uf = new UnionFind(5)
    uf.union(0, 1)
    uf.union(3, 4)
    expect(uf.connected(0, 3)).toBe(false)
  })
  it('全要素を同一集合に: 全ペアがconnected', () => {
    const uf = new UnionFind(4)
    uf.union(0, 1)
    uf.union(1, 2)
    uf.union(2, 3)
    expect(uf.connected(0, 3)).toBe(true)
    expect(uf.connected(1, 3)).toBe(true)
  })
  it('自己同士はconnectedになる', () => {
    const uf = new UnionFind(3)
    expect(uf.connected(0, 0)).toBe(true)
  })
  it('同じunionを複数回呼んでも副作用がない', () => {
    const uf = new UnionFind(3)
    uf.union(0, 1)
    uf.union(0, 1)
    expect(uf.connected(0, 1)).toBe(true)
    expect(uf.connected(1, 2)).toBe(false)
  })
})
