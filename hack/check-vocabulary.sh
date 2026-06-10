#!/usr/bin/env bash
# Protocol P2, enforced at agent level: bold is coinage, and coinage
# requires definition. Every **bold term** in the normative docs must be
# in the built vocabulary (eval/vocabulary.txt, from `make vocab`) or the
# committed baseline of rhetorical emphasis. Agents run this before
# committing doc changes; CI and the drift audit run it after them.
set -u
VOCAB=eval/vocabulary.txt
BASELINE=eval/vocabulary-baseline.txt

norm() { tr '[:upper:]' '[:lower:]' | sed 's/[[:space:]]\+/ /g; s/^ //; s/ $//'; }

coinages=$(grep -ohE '\*\*[^*]+\*\*' docs/*.md docs/assessments/*.md 2>/dev/null \
  | sed 's/\*\*//g' | norm | sort -u)
known=$(cat "$VOCAB" "$BASELINE" 2>/dev/null | norm | sort -u)
violations=$(comm -23 <(echo "$coinages") <(echo "$known"))

if [ -n "$violations" ]; then
  echo "Undefined coinages — define in the lexicon (then 'make vocab'), or baseline in a reviewed commit:"
  echo "$violations" | sed 's/^/  - /'
  exit 1
fi
echo "vocabulary check: all coined terms defined or baselined"
