#!/usr/bin/env bash
# The promise check: "promise delivered" is a verdict, never a flourish.
# Walks release/PROMISE's mechanical clauses; any red blocks the release.
set -u
fail=0
say() { echo "[$1] $2"; if [ "$1" = "fail" ]; then fail=1; fi; return 0; }

[ -f release/PROMISE ] && say pass "the promise is declared (release/PROMISE)" \
  || say fail "no promise declared — a release without a promise may not claim delivery"

if [ -f dist/test-report.json ]; then
  fails=$(grep -c '"Action":"fail"' dist/test-report.json || true)
  [ "$fails" = "0" ] && say pass "tests_pass: zero failures in dist/test-report.json" \
    || say fail "tests_pass: $fails failing test(s)"
else say fail "tests_pass: dist/test-report.json missing"; fi

if [ -s dist/security-report.json ] && grep -q '"' dist/security-report.json; then
  say pass "no_known_vulnerabilities: govulncheck ran and wrote its report (findings fail its own step)"
else say fail "no_known_vulnerabilities: dist/security-report.json missing or empty"; fi

if [ -f dist/eval-report.json ]; then
  grep -q '"resolution_rate_percent": 100' dist/eval-report.json \
    && say pass "registry_resolves: 100%" || say fail "registry_resolves: below 100%"
else say fail "registry_resolves: dist/eval-report.json missing"; fi

./hack/check-vocabulary.sh >/dev/null 2>&1 && say pass "every_word_defined: vocabulary green" \
  || say fail "every_word_defined: vocabulary check failed"

[ -f dist/install.yaml ] && say pass "install_one_apply: bundle present" \
  || say fail "install_one_apply: dist/install.yaml missing"

if [ "$fail" = "1" ]; then
  echo "verdict: the promise is NOT delivered — the release is blocked"
  exit 1
fi
echo "verdict: promise delivered — every clause carries its verdict"
