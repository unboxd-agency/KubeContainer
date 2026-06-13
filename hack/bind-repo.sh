#!/bin/sh
# Bind any repository to a book — one URL in, one book out:
#   ./hack/bind-repo.sh https://github.com/org/repo [out-dir]
# Content unaltered; the story is the repo's own.
set -e
URL="$1"
[ -n "$URL" ] || { echo "usage: bind-repo.sh <git-url> [out-dir]" >&2; exit 2; }
NAME=$(basename "$URL" .git)
OUT="${2:-books/$NAME}"
TMP=$(mktemp -d)
git clone -q --depth 1 "$URL" "$TMP/$NAME"
go run ./cmd/bookbinder -repo "$TMP/$NAME" -title "$NAME" -out "$OUT"
rm -rf "$TMP"
