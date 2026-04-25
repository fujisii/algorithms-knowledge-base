import { describe, it, expect } from 'vitest'
import { knapsack } from './knapsack-01'

describe('knapsack', () => {
  it('正常系: 最適な組み合わせを選ぶ', () => {
    const items = [
      { weight: 2, value: 6 },
      { weight: 2, value: 10 },
      { weight: 3, value: 12 },
    ]
    expect(knapsack(items, 5)).toBe(22)
  })
  it('正常系: 全アイテムが容量内に収まる', () => {
    const items = [
      { weight: 1, value: 1 },
      { weight: 2, value: 2 },
      { weight: 3, value: 3 },
    ]
    expect(knapsack(items, 10)).toBe(6)
  })
  it('境界: 空アイテムリスト', () => {
    expect(knapsack([], 10)).toBe(0)
  })
  it('境界: capacity=0', () => {
    const items = [{ weight: 1, value: 10 }]
    expect(knapsack(items, 0)).toBe(0)
  })
  it('境界: 全アイテムが容量超過', () => {
    const items = [
      { weight: 5, value: 10 },
      { weight: 6, value: 20 },
    ]
    expect(knapsack(items, 4)).toBe(0)
  })
  it('境界: 単一アイテムが丁度容量と等しい', () => {
    const items = [{ weight: 3, value: 15 }]
    expect(knapsack(items, 3)).toBe(15)
  })
  it('境界: weight=0 のアイテムは容量 0 でも選択可能', () => {
    const items = [{ weight: 0, value: 5 }]
    expect(knapsack(items, 0)).toBe(5)
  })
  it('エラー: capacity < 0 は RangeError を投げる', () => {
    expect(() => knapsack([], -1)).toThrow(RangeError)
  })
  it('エラー: weight < 0 のアイテムは RangeError を投げる', () => {
    expect(() => knapsack([{ weight: -1, value: 5 }], 10)).toThrow(RangeError)
  })
})
