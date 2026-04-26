@CLAUDE.local.md

# algorithms-knowledge-base AI 行動指針

詳細な品質基準は `docs/algorithm-guide.md` を参照すること。

## アルゴリズム追加の手順

1. `scripts/scaffold.sh <algorithm-name>` を実行してディレクトリを生成する
2. 生成された README.md の `{{ALGORITHM_NAME_JA}}` と `{{DIFFICULTY}}` を置換する
3. `algorithms/[name]/typescript/[kebab].ts` に実装を作成する
4. `algorithms/[name]/go/[snake].go` に実装を作成する
5. TypeScript テスト・Go テストを作成する（テストケース仕様は `docs/algorithm-guide.md` 参照）
6. README.md の全セクションを記述する（各セクションの指針は `docs/algorithm-guide.md` 参照）
7. `npx vitest run algorithms/[name]/typescript/` でテストが通ることを確認する
8. `go test ./algorithms/[name]/go/...` でテストが通ることを確認する
9. `scripts/lint-readme.sh` で README の完全性を確認する

## 品質の原則

- 実装は可読性優先（最適化より読みやすさ）
- TypeScript: `any` 禁止・strict mode 準拠
- Go: `package main`・エクスポート関数は PascalCase
- テストはテーブルテスト形式（Go）/ `describe-it` 形式（TypeScript）
- README の「参考文献」は一次ソース（教科書・論文・公式ドキュメント）を必ず1件以上含める
