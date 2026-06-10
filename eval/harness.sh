#!/usr/bin/env bash
# Evaluation registry harness: runs every task's world-test and emits an
# evidence record. The registry (eval/corpus) is append-only; a report is
# provenance, not opinion.
set -u
REPORT="${1:-dist/eval-report.json}"
mkdir -p "$(dirname "$REPORT")"

commit=$(git rev-parse HEAD)
when=$(date -u +%Y-%m-%dT%H:%M:%SZ)
total=0; passed=0; entries=""

for task in eval/corpus/*.yaml; do
  id=$(grep -m1 '^id:' "$task" | awk '{print $2}')
  test_cmd=$(awk '/^world_test: >-$/{flag=1;next} flag&&/^[a-z_]+:/{flag=0} flag{printf "%s ",$0}' "$task")
  [ -z "$test_cmd" ] && test_cmd=$(grep -m1 '^world_test:' "$task" | cut -d' ' -f2-)
  total=$((total+1))
  if eval "$test_cmd" >/dev/null 2>&1; then
    result="pass"; passed=$((passed+1))
  else
    result="fail"
  fi
  echo "[$result] $id  ($task)"
  entries="$entries{\"id\":\"$id\",\"task\":\"$task\",\"result\":\"$result\"},"
done

rate=$(awk "BEGIN{printf \"%.2f\", ($passed/$total)*100}")
cat > "$REPORT" <<JSON
{
  "registry": "eval/corpus",
  "commit": "$commit",
  "evaluated_at": "$when",
  "tasks_total": $total,
  "tasks_passed": $passed,
  "resolution_rate_percent": $rate,
  "results": [${entries%,}]
}
JSON
echo "---"
echo "resolution: $passed/$total ($rate%) -> $REPORT"
[ "$passed" -eq "$total" ]
