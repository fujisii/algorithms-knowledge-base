#!/bin/bash

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ALGORITHMS_DIR="$REPO_ROOT/algorithms"

REQUIRED_SECTIONS=(
  "概要"
  "アルゴリズムの直感"
  "自分の解釈"
  "計算量"
  "Example"
  "実装メモ"
  "参考文献"
)

FAILED=()

for algo_dir in "$ALGORITHMS_DIR"/*/; do
  [ -d "$algo_dir" ] || continue
  name="$(basename "$algo_dir")"
  readme="$algo_dir/README.md"

  if [ ! -f "$readme" ]; then
    echo "MISSING README: $name"
    FAILED+=("$name")
    continue
  fi

  errors=()

  if ! grep -q '\*\*難易度\*\*' "$readme"; then
    errors+=("**難易度** フィールドが見つかりません")
  fi

  for section in "${REQUIRED_SECTIONS[@]}"; do
    if ! grep -q "^## $section" "$readme"; then
      errors+=("セクション '## $section' が見つかりません")
    fi
  done

  if grep -q '{{' "$readme"; then
    placeholders="$(grep -o '{{[^}]*}}' "$readme" | sort -u | tr '\n' ' ')"
    errors+=("未置換プレースホルダーが残っています: $placeholders")
  fi

  if grep -q "^## 参考文献" "$readme"; then
    if grep -A 10 "^## 参考文献" "$readme" | grep -q "TODO"; then
      errors+=("参考文献セクションに TODO が残っています")
    fi
  fi

  if [ "${#errors[@]}" -gt 0 ]; then
    echo "FAILED: $name"
    for e in "${errors[@]}"; do
      echo "  - $e"
    done
    FAILED+=("$name")
  fi
done

echo ""
if [ "${#FAILED[@]}" -eq 0 ]; then
  echo "All READMEs passed."
  exit 0
else
  echo "FAILED: ${#FAILED[@]} algorithm(s)"
  exit 1
fi
