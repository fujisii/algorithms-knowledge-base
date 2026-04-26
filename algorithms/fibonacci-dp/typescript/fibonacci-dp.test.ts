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
  it('正常系: n=50 は正確な値を返す', () => {
    expect(fibonacci(50)).toBe(12586269025)
  })
  it('境界: n=78 は Number.MAX_SAFE_INTEGER 以内で正確', () => {
    // F(79) 以降は IEEE 754 倍精度の精度限界(2^53-1)を超えるため結果が不正確になる
    expect(fibonacci(78)).toBe(8944394323791464)
  })
  it('境界: n=79 は精度限界を超えるため真値と一致しない（ドキュメント目的）', () => {
    // 真値 F(79) = 14472334024676221 だが IEEE 754 倍精度では 14472334024676220 に丸められる
    // JS の number では両値が同一 double になるため BigInt で比較する
    const result = fibonacci(79)
    expect(BigInt(result)).not.toBe(14472334024676221n)
  })
  it('エラー: n < 0 は RangeError を投げる', () => {
    expect(() => fibonacci(-1)).toThrow(RangeError)
  })
})
