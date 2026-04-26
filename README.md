# algorithms-knowledge-base

アルゴリズムを体系的に学習・整理する知識資産リポジトリ。TypeScript と Go による実行可能な実装・テスト、「再実装できる解説」を含む README を標準フォーマットで管理する。

## アルゴリズム一覧

| Algorithm | 日本語名 | 難易度 | 概要 |
|-----------|---------|--------|------|
| [Binary Search](./algorithms/binary-search/) | 二分探索 | Easy | ソート済み配列を半分に絞り込みながら目標値を探索する |
| [Bubble Sort](./algorithms/bubble-sort/) | バブルソート | Easy | 隣接要素を比較・交換しながら最大値を末尾に移動し繰り返す |
| [Quick Sort](./algorithms/quick-sort/) | クイックソート | Medium | ピボットを選択して小・大に分割し再帰的にソートする |
| [Merge Sort](./algorithms/merge-sort/) | マージソート | Medium | 配列を半分に分割し再帰的にソートしてマージする安定ソート |
| [BFS](./algorithms/bfs/) | 幅優先探索 | Medium | キューを使ってグラフを始点から近い順に幅優先で探索する |
| [DFS](./algorithms/dfs/) | 深さ優先探索 | Medium | スタック（または再帰）を使ってグラフを深さ優先で探索する |
| [Dijkstra](./algorithms/dijkstra/) | ダイクストラ法 | Hard | 負辺なし重み付きグラフの単一始点最短経路を貪欲法で求める |
| [Union-Find](./algorithms/union-find/) | 素集合データ構造 | Medium | パス圧縮とランクで最適化した集合の合併・判定データ構造 |
| [Fibonacci DP](./algorithms/fibonacci-dp/) | フィボナッチ数列（動的計画法） | Easy | DPテーブルで重複計算を排除しフィボナッチ数列を線形時間で求める |
| [0/1 Knapsack](./algorithms/knapsack-01/) | 0/1 ナップサック問題 | Medium | 2次元DPで各アイテムを取る・取らないを管理し最大価値を求める |

## 使い方

```bash
# TypeScript 実行
npx tsx algorithms/binary-search/typescript/run.ts

# TypeScript テスト（単体）
npx vitest run algorithms/binary-search/typescript/

# TypeScript テスト（全体）
npm test

# Go 実行
go run ./algorithms/binary-search/go/

# Go テスト
go test ./algorithms/binary-search/go/...

# 全テスト一括
./scripts/test-all.sh

# README 完全性検証
./scripts/lint-readme.sh

# 新アルゴリズム追加
./scripts/scaffold.sh <algorithm-name>
```

## ディレクトリ構成

```
algorithms/[name]/
├── README.md          # 再実装可能な解説（8セクション必須）
├── typescript/
│   ├── [kebab].ts     # 実装
│   ├── [kebab].test.ts # テスト
│   └── run.ts         # 実行エントリ
└── go/
    ├── [snake].go     # 実装
    ├── [snake]_test.go # テスト
    └── run.go         # 実行エントリ
```
