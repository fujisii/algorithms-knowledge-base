# Binary Search / 二分探索

**難易度**: Easy

## 概要

ソート済み配列に対して、探索範囲を毎ステップで半分に絞り込むことで目的の要素を効率的に見つける探索アルゴリズム。未整列配列には適用できない。返り値はインデックス値で、未検出時は `-1` を返す。

## アルゴリズムの直感

左端・右端のポインタで囲んだ範囲の中央値を目標値と比較する。目標値が中央値より大きければ左半分を捨て、小さければ右半分を捨てる。この操作を繰り返すと、探索範囲は毎回半分になる。

Cormen et al. (CLRS) では「分割統治の簡略形」として説明されており、「同じ問題をより小さなサイズで再帰的に解く」という考え方に基づく。

## 自分の解釈

最初に迷いやすいのは `left <= right` と `left < right` の違い。`<=` にしないと、配列が1要素のとき（`left == right`）に中央を一度も見ずに終わる。また、中央インデックスを `(left + right) / 2` で計算すると整数オーバーフローが起きるケースがある（Java/C 等では問題になる）が、JavaScript/TypeScript の数値はすべて浮動小数点なので実用上は問題ない。

## 計算量

- 時間: O(log n)
- 空間: O(1)

## Example

```
入力: arr = [1, 3, 5, 7, 9], target = 5

Step 1: left=0, right=4, mid=2, arr[2]=5 → target と一致 → return 2

入力: arr = [1, 3, 5, 7, 9], target = 6

Step 1: left=0, right=4, mid=2, arr[2]=5 < 6 → left=3
Step 2: left=3, right=4, mid=3, arr[3]=7 > 6 → right=2
Step 3: left=3 > right=2 → ループ終了 → return -1
```

## 実装メモ

- 配列のインプレース変更をしないため、引数は `readonly number[]` で受け取る
- 「中央を計算して比較 → 範囲を絞る」の2ステップが1ループ。ループ不変条件: 目標値が存在するなら `arr[left..right]` の中にある

## 参考文献

- Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022). *Introduction to Algorithms* (4th ed.), Section 2.3. MIT Press.
