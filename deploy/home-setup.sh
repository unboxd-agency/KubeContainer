#!/usr/bin/env bash
# The home, in one script — idempotent, declarative, verdict-gated.
# Clone the repo on the VPS and run:  sudo bash deploy/home-setup.sh
# Every decision is a variable below (Config As Code); run it twice and
# the second run changes nothing that already converged.
set -u

### The declaration (edit here, nothing below) ###############################
GITLAB_HOST="git.unboxd.cloud"
GITLAB_VERSION="19.0.2-ce.0"     # the pin, recorded
GITLAB_HTTP_PORT="8081"          # GitLab's nginx, behind the front proxy
TRAEFIK_POLICY="route"           # route = keep Traefik on 80/443 and route through it
                                 # evict = stop/disable standalone Traefik, GitLab takes 80
INSTALL_K0S="yes"                # the runtime rung, same box
KC_BUNDLE="https://github.com/unboxd-cloud/KubeContainer/releases/download/v0.2.0/install.yaml"
##############################################################################

pass=0; fail=0
say(){ echo "[$1] $2"; if [ "$1" = pass ]; then pass=$((pass+1)); else fail=$((fail+1)); fi; }

echo "== 1. ground =="
apt-get update -qq && apt-get install -y -qq curl ca-certificates tzdata perl >/dev/null \
  && say pass "prerequisites" || say fail "prerequisites"

echo "== 2. gitlab repo (noble channel on newer ubuntu, recorded off-label) =="
if ! apt-cache madison gitlab-ce 2>/dev/null | grep -q gitlab-ce; then
  curl -s https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh \
    | os=ubuntu dist=noble bash >/dev/null && say pass "gitlab repo added" || say fail "gitlab repo"
else say pass "gitlab repo present"; fi

echo "== 3. traefik seating: $TRAEFIK_POLICY =="
TRAEFIK_OWNS_80="no"
ss -tlnp 2>/dev/null | grep -E ':80 ' | grep -q traefik && TRAEFIK_OWNS_80="yes"
if [ "$TRAEFIK_OWNS_80" = yes ] && [ "$TRAEFIK_POLICY" = evict ]; then
  systemctl disable --now traefik 2>/dev/null; docker rm -f traefik 2>/dev/null
  say pass "traefik evicted (one service per port)"
fi

echo "== 4. gitlab install + certless config (cert-manager issues later) =="
if ! dpkg -l gitlab-ce 2>/dev/null | grep -q '^ii'; then
  EXTERNAL_URL="http://$GITLAB_HOST" apt-get install -y gitlab-ce="$GITLAB_VERSION" >/dev/null 2>&1
fi
GLRB=/etc/gitlab/gitlab.rb
grep -q "^external_url 'http://$GITLAB_HOST'" $GLRB 2>/dev/null \
  || sed -i "s|^external_url .*|external_url 'http://$GITLAB_HOST'|" $GLRB
grep -q "^letsencrypt\['enable'\] = false" $GLRB \
  || echo "letsencrypt['enable'] = false" >> $GLRB
if [ "$TRAEFIK_POLICY" = route ] && [ "$TRAEFIK_OWNS_80" = yes ]; then
  grep -q "nginx\['listen_port'\] = $GITLAB_HTTP_PORT" $GLRB \
    || printf "nginx['listen_port'] = %s\nnginx['listen_https'] = false\n" "$GITLAB_HTTP_PORT" >> $GLRB
fi
gitlab-ctl reconfigure >/dev/null 2>&1 && say pass "gitlab reconfigured" || say fail "gitlab reconfigure (run: sudo gitlab-ctl reconfigure)"
dpkg --configure -a >/dev/null 2>&1
gitlab-ctl status >/dev/null 2>&1 && say pass "gitlab services up" || say fail "gitlab services"

echo "== 5. runtime (k0s): $INSTALL_K0S =="
if [ "$INSTALL_K0S" = yes ]; then
  if ! command -v k0s >/dev/null; then
    curl -sSLf https://get.k0s.sh | sh >/dev/null 2>&1
    k0s install controller --single 2>/dev/null; k0s start; sleep 20
  fi
  k0s kubectl get nodes 2>/dev/null | grep -q Ready && say pass "k0s node Ready" || say fail "k0s node not Ready yet (re-run in a minute)"
  k0s kubectl apply -f "$KC_BUNDLE" >/dev/null 2>&1 && say pass "KubeContainer v0.2.0 applied" || say fail "operator apply"
fi

echo "== verdicts: $pass pass, $fail fail =="
echo "next: initial password -> cat /etc/gitlab/initial_root_password"
[ "$TRAEFIK_POLICY" = route ] && [ "$TRAEFIK_OWNS_80" = yes ] && \
  echo "next: point Traefik at 127.0.0.1:$GITLAB_HTTP_PORT for host $GITLAB_HOST (its config location depends on who installed it: docker label or /etc/traefik)"
echo "next: cert-manager (DNS-01) into k0s issues https for everything — flip external_url then"
[ "$fail" = 0 ]
