#!/bin/bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

if [ $# -eq 0 ]; then
  echo "Usage: $(basename "$0") <algorithm-name>"
  echo "  Example: $(basename "$0") binary-search"
  exit 1
fi

NAME="$1"

if [[ ! "$NAME" =~ ^[a-z0-9-]+$ ]]; then
  echo "Error: name must be kebab-case (lowercase letters, digits, and hyphens only)"
  echo "  Example: $(basename "$0") binary-search"
  exit 1
fi

ALGO_DIR="$REPO_ROOT/algorithms/$NAME"
TEMPLATES_DIR="$REPO_ROOT/templates"

if [ -d "$ALGO_DIR" ]; then
  echo "Error: $ALGO_DIR already exists"
  exit 1
fi

# kebab-case to snake_case
SNAKE="$(echo "$NAME" | tr '-' '_')"

# kebab-case to PascalCase
PASCAL="$(echo "$NAME" | awk 'BEGIN{FS="-"; OFS=""} {for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2); print}')"

# kebab-case to camelCase
FIRST_LOWER="$(echo "${PASCAL:0:1}" | tr '[:upper:]' '[:lower:]')"
CAMEL="${FIRST_LOWER}${PASCAL:1}"

mkdir -p "$ALGO_DIR/typescript"
mkdir -p "$ALGO_DIR/go"

# copy and rename files
cp "$TEMPLATES_DIR/README.md" "$ALGO_DIR/README.md"
cp "$TEMPLATES_DIR/typescript/algorithm.ts" "$ALGO_DIR/typescript/$NAME.ts"
cp "$TEMPLATES_DIR/typescript/algorithm.test.ts" "$ALGO_DIR/typescript/$NAME.test.ts"
cp "$TEMPLATES_DIR/typescript/run.ts" "$ALGO_DIR/typescript/run.ts"
cp "$TEMPLATES_DIR/go/algorithm.go" "$ALGO_DIR/go/${SNAKE}.go"
cp "$TEMPLATES_DIR/go/algorithm_test.go" "$ALGO_DIR/go/${SNAKE}_test.go"
cp "$TEMPLATES_DIR/go/run.go" "$ALGO_DIR/go/run.go"

# replace placeholders in all copied files (POSIX sed — works on macOS and Linux)
find "$ALGO_DIR" -type f | while read -r file; do
  tmp=$(mktemp)
  sed -e "s/{{ALGORITHM_NAME_PASCAL}}/$PASCAL/g" \
      -e "s/{{ALGORITHM_NAME_CAMEL}}/$CAMEL/g" \
      -e "s/{{ALGORITHM_NAME_SNAKE}}/$SNAKE/g" \
      -e "s/{{ALGORITHM_NAME}}/$NAME/g" \
      "$file" > "$tmp" && mv "$tmp" "$file"
done

# warn if any unreplaced placeholders remain
remaining="$(grep -rl '{{' "$ALGO_DIR" 2>/dev/null || true)"
if [ -n "$remaining" ]; then
  echo "Warning: unreplaced placeholders remain in:"
  echo "$remaining" | while read -r f; do
    echo "  $f"
    grep -n '{{' "$f"
  done
fi

echo "Scaffolded: $ALGO_DIR"
