import { describe, it, expect } from 'vitest'
import { binarySearch } from './binary-search'

describe('binarySearch', () => {
  it('正常系: 中央の要素を返す', () => {
    expect(binarySearch([1, 3, 5, 7, 9], 5)).toBe(2)
  })
  it('正常系: 左端の要素を返す', () => {
    expect(binarySearch([1, 3, 5, 7, 9], 1)).toBe(0)
  })
  it('正常系: 右端の要素を返す', () => {
    expect(binarySearch([1, 3, 5, 7, 9], 9)).toBe(4)
  })
  it('境界: 空配列は -1 を返す', () => {
    expect(binarySearch([], 1)).toBe(-1)
  })
  it('境界: 単一要素で一致', () => {
    expect(binarySearch([5], 5)).toBe(0)
  })
  it('境界: 単一要素で不一致', () => {
    expect(binarySearch([5], 3)).toBe(-1)
  })
  it('異常系: 未検出は -1 を返す', () => {
    expect(binarySearch([1, 3, 5, 7, 9], 4)).toBe(-1)
  })
  it('正常系: 重複要素あり（有効なインデックスのいずれかを返す）', () => {
    const idx = binarySearch([1, 3, 3, 3, 5], 3)
    expect(idx).toBeGreaterThanOrEqual(1)
    expect(idx).toBeLessThanOrEqual(3)
  })
})
