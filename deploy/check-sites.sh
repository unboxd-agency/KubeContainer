#!/bin/sh
# End-to-end check of every public face: HTTP status, TLS validity, and
# a content marker proving the correct site is served. Run from anywhere
# with internet:  sh deploy/check-sites.sh
set -u
check() {
  url="$1"; want="$2"
  body=$(curl -sS --max-time 15 -w '\n%{http_code}' "$url" 2>/tmp/cs_err)
  rc=$?
  code=$(printf '%s' "$body" | tail -n1)
  html=$(printf '%s' "$body" | sed '$d')
  if [ $rc -ne 0 ]; then tls=$(grep -qi 'certificate' /tmp/cs_err && echo "CERT-INVALID" || echo "CONN-FAIL"); else tls="valid"; fi
  if printf '%s' "$html" | grep -qi "$want"; then content="ok"; else content="MISSING[$want]"; fi
  printf '%-34s  http=%-3s  tls=%-12s  %s\n' "$url" "${code:-?}" "$tls" "$content"
}
echo "=== Fabric estate e2e ==="
check https://www.openautonomyx.com       "Enterprise Work"
check https://platform.openautonomyx.com  "Arithmetic Platform"
check https://fabric.openautonomyx.com    "Fabric Browser"
check https://kubecontainer.xyz           "KubeContainer"
check https://registry.agennext.com       "Agent Registry"
check https://agennext.space              "FileFabric"
echo "(http=200, tls=valid, content=ok on every line = the estate is green)"
