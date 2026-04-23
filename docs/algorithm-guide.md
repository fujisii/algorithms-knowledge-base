# アルゴリズム追加品質ガイド

## 難易度判断基準

| 難易度 | 定義 | 例 |
|--------|------|-----|
| Easy | 単一ループまたは再帰で実装可能。データ構造の選択が自明 | Binary Search, Bubble Sort, Fibonacci DP |
| Medium | 複数のデータ構造を組み合わせる・再帰 + 状態管理が必要 | Quick Sort, Merge Sort, BFS, DFS, Union-Find, 0/1 Knapsack |
| Hard | グラフアルゴリズム・動的計画法・複雑な状態空間を扱う | Dijkstra |

## README 各セクションの書き方

### 難易度
`**難易度**: Easy / Medium / Hard` の形式で記述する。行頭から始める。

### 概要
何をするアルゴリズムか、2〜4文で説明する。入力・出力・用途を含める。

### アルゴリズムの直感
「なぜこの手順でうまくいくか」の直感的な説明。図・比喩を使ってよい。一次ソース（教科書・論文）からの根拠を含める。

### 自分の解釈
初めて理解したときに何が難しかったか、どう考えると腑に落ちたかの個人的な記録。要約ではなく思考プロセスを残す。

### 計算量
- 時間: O(...) — 最良・平均・最悪を区別する（必要な場合）
- 空間: O(...) — 補助空間のみ or 入力含む場合は明記する

### Example の trace 形式
入力 → 各ステップで変数がどう変化したか → 出力 の順で記述する。表形式推奨。

```
入力: arr = [1, 3, 5, 7, 9], target = 5

Step 1: left=0, right=4, mid=2, arr[2]=5 → 一致
出力: 2
```

### 実装メモ
実装上の注意点・落とし穴・ハマりやすい点を記述する。テンプレートの TODO を削除して実際の内容を書く。

### 参考文献
一次ソース（教科書・論文・公式ドキュメント）を最低1件記載する。ソース種別によって書式と URL の要否が異なる（下記「参考文献の記載形式」を参照）。

## 一次ソース選定指針

優先順位（高い順）:
1. 教科書（CLRS: Introduction to Algorithms, Sedgewick: Algorithms 等）
2. 論文（ACM DL, arXiv 等）
3. 言語・標準ライブラリの公式ドキュメント
4. 信頼性の高い技術記事（著者・所属が明確なもの）

避けるべき:
- Wikipedia 単独引用（一次ソースが明示されていない場合）
- 個人ブログ（一次ソースが不明なもの）

## 参考文献の記載形式

ソース種別ごとに書式と URL の要否が異なる。

### 教科書・書籍（URL なし）

出版社ページや購入リンクは参照先として意味がないため URL は付けない。著者・タイトル・版・該当箇所（章・節・ページ）・出版社を記載する。

```
- 著者名. (出版年). *タイトル* (版), Section X.X / Chapter X / p.XXX. 出版社.
```

**例:**
```
- Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022). *Introduction to Algorithms* (4th ed.), Section 2.1. MIT Press.
- Sedgewick, R., & Wayne, K. (2011). *Algorithms* (4th ed.), Section 2.1. Addison-Wesley Professional.
```

### 論文（DOI URL あり）

DOI は永続識別子であり参照先として安定しているため、DOI URL を付ける。

```
- 著者名. (出版年). "論文タイトル". *ジャーナル名*, vol(issue), pp.XXX–XXX. https://doi.org/xxxxx
```

**例:**
```
- Dijkstra, E. W. (1959). "A note on two problems in connexion with graphs". *Numerische Mathematik*, 1(1), pp.269–271. https://doi.org/10.1007/BF01386390
```

### 公式ドキュメント（URL あり）

言語仕様・標準ライブラリ・RFC 等の公式ドキュメントは URL を付ける。アクセス年も記載する。

```
- 組織名. (アクセス年). "ページタイトル". URL
```

**例:**
```
- Go Authors. (2024). "Package sort — Go standard library". https://pkg.go.dev/sort
- IETF. (2024). "RFC 4122: A Universally Unique IDentifier (UUID) URN Namespace". https://www.rfc-editor.org/rfc/rfc4122
```

## テストケース設計方針

要件書 要件 6 の「アルゴリズム別テストケース方針」を参照。

| カテゴリ | 対象アルゴリズム | 必須テストケース |
|----------|-----------------|----------------|
| 探索系 | Binary Search | 正常系（中央・左端・右端）、空配列、単一要素一致、単一要素不一致、未検出 |
| ソート系 | Bubble / Quick / Merge Sort | 正常系、空配列、単一要素、既ソート済み、逆順、重複あり |
| グラフ系 | BFS / DFS | 正常系、単一ノード、到達不可能なノードが含まれる、閉路あり |
| 重み付きグラフ | Dijkstra | 正常系、単一ノード（距離0）、到達不可能なノード（Infinity） |
| 集合系 | Union-Find | 全要素独立状態、union 後の connected、全要素同一集合 |
| DP 系 | Fibonacci DP | n=0、n=1、n=10 以上、n < 0（エラー） |
| DP 系 | 0/1 Knapsack | 正常系、空アイテムリスト、capacity=0、全アイテムが容量超過 |

### テスト構造（TypeScript）

```typescript
import { describe, it, expect } from 'vitest'
import { targetFunction } from './algorithm-name'

describe('targetFunction', () => {
  it('正常系: ...', () => {
    expect(targetFunction(...)).toBe(...)
  })
  it('境界: 空配列は -1 を返す', () => {
    expect(targetFunction([], ...)).toBe(-1)
  })
})
```

### テスト構造（Go）

```go
func TestTargetFunction(t *testing.T) {
  tests := []struct {
    name string
    // 入力フィールド
    want // 期待値
  }{
    {"正常系: ...", ..., ...},
    {"境界: 空スライス", ..., ...},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got := TargetFunction(...)
      if got != tt.want {
        t.Errorf("got %v; want %v", got, tt.want)
      }
    })
  }
}
```

## 関数シグネチャ設計指針

- 副作用を持つ関数（配列のインプレース変更等）は関数名または引数名で明示する
- TypeScript: 引数・返り値の型を明示（`any` 禁止）。配列を変更しない関数は `readonly` を付与
- Go: エクスポート関数は PascalCase、内部関数は camelCase

## エラー処理方針

アルゴリズムが前提条件を満たさない入力（負容量・負辺・範囲外インデックスなど）を受け取った場合、両言語で同等の明示的なエラーを発生させる。

| 言語 | 手段 | 例 |
|------|------|----|
| TypeScript | `throw new RangeError("...")` | 負のインデックス・負容量・前提条件違反 |
| Go | `panic(fmt.Sprintf("..."))` | 同上 |

**原則:**
- エラーメッセージには「何が」「期待値と実際値」を含める（例: `capacity must be >= 0, got -1`）
- TypeScript と Go で同じ条件に対して同等のガード節を実装する（非対称禁止）
- アルゴリズムの数学的前提条件（ダイクストラ: 非負辺、ナップサック: 非負容量 等）は必ずコードで検証する
- テストでは正常系と同様に、エラーケースも両言語でカバーする
