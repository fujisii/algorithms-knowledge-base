import { describe, it, expect } from 'vitest'
import { bubbleSort } from './bubble-sort'

describe('bubbleSort', () => {
  it('正常系: 整数配列をソートする', () => {
    expect(bubbleSort([5, 3, 1, 4, 2])).toEqual([1, 2, 3, 4, 5])
  })
  it('境界: 空配列はそのまま返す', () => {
    expect(bubbleSort([])).toEqual([])
  })
  it('境界: 単一要素配列はそのまま返す', () => {
    expect(bubbleSort([42])).toEqual([42])
  })
  it('境界: 既ソート済み配列はそのまま返す', () => {
    expect(bubbleSort([1, 2, 3, 4, 5])).toEqual([1, 2, 3, 4, 5])
  })
  it('境界: 逆順配列を正しくソートする', () => {
    expect(bubbleSort([5, 4, 3, 2, 1])).toEqual([1, 2, 3, 4, 5])
  })
  it('境界: 重複要素を含む配列をソートする', () => {
    expect(bubbleSort([3, 1, 2, 1, 3])).toEqual([1, 1, 2, 3, 3])
  })
  it('境界: 負数を含む配列をソートする', () => {
    expect(bubbleSort([-3, 1, -1, 2, 0])).toEqual([-3, -1, 0, 1, 2])
  })
})
