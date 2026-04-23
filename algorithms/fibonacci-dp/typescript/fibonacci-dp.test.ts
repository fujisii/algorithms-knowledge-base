import { describe, it, expect } from 'vitest'
import { fibonacci } from './fibonacci-dp'

describe('fibonacci', () => {
  it('境界: n=0 は 0 を返す', () => {
    expect(fibonacci(0)).toBe(0)
  })
  it('境界: n=1 は 1 を返す', () => {
    expect(fibonacci(1)).toBe(1)
  })
  it('正常系: n=2 は 1 を返す', () => {
    expect(fibonacci(2)).toBe(1)
  })
  it('正常系: n=10 は 55 を返す', () => {
    expect(fibonacci(10)).toBe(55)
  })
  it('正常系: n=20 は 6765 を返す', () => {
    expect(fibonacci(20)).toBe(6765)
  })
  it('エラー: n < 0 は RangeError を投げる', () => {
    expect(() => fibonacci(-1)).toThrow(RangeError)
  })
})
