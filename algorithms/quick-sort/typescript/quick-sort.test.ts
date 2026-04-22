import { describe, it, expect } from 'vitest'
import { quickSort } from './quick-sort'

describe('quickSort', () => {
  it('正常系: 一般的な配列をソートする', () => {
    expect(quickSort([3, 1, 4, 1, 5, 9, 2, 6])).toEqual([1, 1, 2, 3, 4, 5, 6, 9])
  })
  it('正常系: 既ソート済み配列', () => {
    expect(quickSort([1, 2, 3, 4, 5])).toEqual([1, 2, 3, 4, 5])
  })
  it('正常系: 逆順配列', () => {
    expect(quickSort([5, 4, 3, 2, 1])).toEqual([1, 2, 3, 4, 5])
  })
  it('境界: 空配列', () => {
    expect(quickSort([])).toEqual([])
  })
  it('境界: 単一要素', () => {
    expect(quickSort([42])).toEqual([42])
  })
  it('境界: 重複要素あり', () => {
    expect(quickSort([3, 3, 1, 1, 2, 2])).toEqual([1, 1, 2, 2, 3, 3])
  })
  it('境界: 全要素が同じ', () => {
    expect(quickSort([5, 5, 5, 5])).toEqual([5, 5, 5, 5])
  })
})
