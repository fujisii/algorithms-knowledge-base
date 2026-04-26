#!/bin/bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ALGORITHMS_DIR="$REPO_ROOT/algorithms"

FAILED=()

for algo_dir in "$ALGORITHMS_DIR"/*/; do
  [ -d "$algo_dir" ] || continue
  name="$(basename "$algo_dir")"

  echo "=== $name ==="

  if [ -d "$algo_dir/typescript" ]; then
    if (cd "$REPO_ROOT" && npx vitest run "algorithms/$name/typescript/"); then
      echo "  TS: PASS"
    else
      echo "  TS: FAIL"
      FAILED+=("$name (TS)")
    fi
  fi

  if [ -d "$algo_dir/go" ]; then
    if (cd "$REPO_ROOT" && go test "./algorithms/$name/go/..."); then
      echo "  Go: PASS"
    else
      echo "  Go: FAIL"
      FAILED+=("$name (Go)")
    fi
  fi
done

echo ""
echo "=== Summary ==="
if [ "${#FAILED[@]}" -eq 0 ]; then
  echo "All algorithms passed."
  exit 0
else
  echo "FAILED:"
  for f in "${FAILED[@]}"; do
    echo "  - $f"
  done
  exit 1
fi
