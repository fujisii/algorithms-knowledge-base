# Union-Find / 素集合データ構造

**難易度**: Medium

## 概要

Union-Find（Disjoint Set Union、DSU とも呼ぶ）は、要素の集合をグループ管理するデータ構造である。「2要素を同じグループに統合する（Union）」と「ある要素がどのグループに属するか調べる（Find）」の2操作を効率的に提供する。グラフの連結性判定やクラスタリングに広く使われる。

## アルゴリズムの直感

各グループを木構造で表現し、木の根（root）をそのグループの代表とする。`find(x)` は x から根まで親を辿り、`union(x, y)` は x と y の根を互いの親としてつなぐ。

**パス圧縮（Path Compression）**: `find` で根を辿る途中のすべてのノードを直接根に接続し直す。次回以降の `find` が O(1) で完了するようになる。

```
before find(4):  1 → 2 → 3 → root
after find(4):   1 → root
                 2 → root
                 3 → root
```

**Union by Rank**: 浅い木を深い木の子につなぐことで、木の高さが O(log n) 以上に伸びるのを防ぐ。パス圧縮と組み合わせると、操作あたりの償却計算量がほぼ O(1)（逆アッカーマン関数 α(n)）になる。

Cormen et al. (2022) は「パス圧縮と Union by Rank の両方を使うことで、m 回の操作の合計時間が O(m α(n)) になる」ことを証明している（α(n) は実用上常に 4 以下）。

## 自分の解釈

最初は「なぜ木の高さをランクで管理するのか」が分からなかった。単純に「大きい木の根に小さい木をつなぐ」だと最悪ケースでも O(log n) にできるが、パス圧縮と組み合わせると rank が実際の木の高さと乖離することがある（パス圧縮で木が扁平化するため）。それでも Union by Rank は「最初の tree height の上限」として機能し、操作全体の償却計算量の理論的保証を与える役割を担っている。

実装上の注意点として、パス圧縮の実装には「パス上の全ノードを根につなぐ完全パス圧縮」と「2段階の圧縮（path halving）」がある。本実装は再帰による完全パス圧縮を採用した（実装がシンプルなため）。

## 計算量

- 時間: O(α(n))（ほぼ O(1)、逆アッカーマン関数）、初期化は O(n)
- 空間: O(n)（parent と rank の配列）

## Example

```
size=5 で初期化: parent=[0,1,2,3,4], rank=[0,0,0,0,0]

union(0, 1):
  find(0)=0, find(1)=1
  rank[0]==rank[1] なので parent[1]=0, rank[0]++
  parent=[0,0,2,3,4], rank=[1,0,0,0,0]

union(1, 2):
  find(1)=0, find(2)=2   ← パス圧縮: parent[1] は既に 0
  rank[0]>rank[2] なので parent[2]=0
  parent=[0,0,0,3,4], rank=[1,0,0,0,0]

union(3, 4):
  find(3)=3, find(4)=4
  rank[3]==rank[4] なので parent[4]=3, rank[3]++
  parent=[0,0,0,3,3], rank=[1,0,0,1,0]

connected(0, 2): find(0)=0, find(2)=0 → true
connected(0, 3): find(0)=0, find(3)=3 → false
```

## 実装メモ

- `find` の再帰実装は深い木でスタックオーバーフローになりうる。要素数が非常に多い場合は反復実装を検討する
- `union` で同じ根を持つ場合は早期リターンすることで rank の不整合を防ぐ
- Go では `*UnionFind` をレシーバとして使わないと `Find` がパス圧縮の結果を呼び出し元に反映できないため、必ずポインタレシーバを使う

## 参考文献

- Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022). *Introduction to Algorithms* (4th ed.), Chapter 21 Data Structures for Disjoint Sets. MIT Press.
