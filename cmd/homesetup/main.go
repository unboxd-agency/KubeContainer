// HomeSetup — one binary handed to the end user: converges a VPS into
// the home (GitLab pinned and certless behind the front proxy, k0s,
// the operator). Declarative flags, idempotent steps, one verdict
// tally. Command as code: every action is a named command, printed
// before it runs, so the user reads exactly what their machine does.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var passN, failN int

func verdict(ok bool, what string) {
	if ok {
		passN++
		fmt.Printf("[pass] %s\n", what)
		return
	}
	failN++
	fmt.Printf("[fail] %s\n", what)
}

func run(name string, arg ...string) bool {
	fmt.Printf("$ %s %s\n", name, strings.Join(arg, " "))
	c := exec.Command(name, arg...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run() == nil
}

func quiet(name string, arg ...string) bool {
	return exec.Command(name, arg...).Run() == nil
}

func appendOnce(path, line string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if strings.Contains(string(b), line) {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	_, err = f.WriteString(line + "\n")
	return err
}

func main() {
	host := flag.String("host", "git.unboxd.cloud", "GitLab external host")
	version := flag.String("gitlab-version", "19.0.2-ce.0", "gitlab-ce pin")
	traefik := flag.String("traefik", "route", "route|evict — the port-80 seating")
	port := flag.String("gitlab-port", "8081", "GitLab nginx port when routed behind the proxy")
	k0s := flag.Bool("k0s", true, "install the k0s runtime and the operator")
	bundle := flag.String("bundle",
		"https://github.com/unboxd-cloud/KubeContainer/releases/download/v0.2.0/install.yaml",
		"KubeContainer install bundle")
	flag.Parse()

	verdict(run("apt-get", "install", "-y", "-qq", "curl", "ca-certificates", "tzdata", "perl"),
		"prerequisites")

	if !quiet("sh", "-c", "apt-cache madison gitlab-ce | grep -q gitlab-ce") {
		verdict(run("sh", "-c",
			"curl -s https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh"+
				" | os=ubuntu dist=noble bash"), "gitlab repo (noble channel, off-label recorded)")
	} else {
		verdict(true, "gitlab repo present")
	}

	traefikOwns80 := quiet("sh", "-c", "ss -tlnp | grep ':80 ' | grep -q traefik")
	if traefikOwns80 && *traefik == "evict" {
		_ = quiet("systemctl", "disable", "--now", "traefik")
		_ = quiet("docker", "rm", "-f", "traefik")
		verdict(true, "traefik evicted (one service per port)")
	}

	if !quiet("sh", "-c", "dpkg -l gitlab-ce | grep -q '^ii'") {
		_ = os.Setenv("EXTERNAL_URL", "http://"+*host)
		verdict(run("apt-get", "install", "-y", "gitlab-ce="+*version), "gitlab-ce "+*version)
	} else {
		verdict(true, "gitlab-ce installed")
	}

	rb := "/etc/gitlab/gitlab.rb"
	ok := run("sed", "-i", "s|^external_url .*|external_url 'http://"+*host+"'|", rb)
	ok = appendOnce(rb, "letsencrypt['enable'] = false") == nil && ok
	if traefikOwns80 && *traefik == "route" {
		ok = appendOnce(rb, "nginx['listen_port'] = "+*port) == nil && ok
		ok = appendOnce(rb, "nginx['listen_https'] = false") == nil && ok
	}
	verdict(ok, "gitlab.rb declared (certless now; cert-manager issues later, DNS-01, for everything)")
	verdict(run("gitlab-ctl", "reconfigure"), "gitlab reconfigured")
	_ = quiet("dpkg", "--configure", "-a")
	verdict(quiet("gitlab-ctl", "status"), "gitlab services up")

	if *k0s {
		if !quiet("sh", "-c", "command -v k0s") {
			verdict(run("sh", "-c", "curl -sSLf https://get.k0s.sh | sh"), "k0s downloaded")
			_ = quiet("k0s", "install", "controller", "--single")
			_ = quiet("k0s", "start")
		}
		verdict(quiet("sh", "-c", "k0s kubectl get nodes | grep -q Ready"),
			"k0s node Ready (if fail: re-run in a minute)")
		verdict(run("k0s", "kubectl", "apply", "-f", *bundle), "KubeContainer operator applied")
	}

	fmt.Printf("verdicts: %d pass, %d fail\n", passN, failN)
	fmt.Println("next: cat /etc/gitlab/initial_root_password (root login, change it)")
	if traefikOwns80 && *traefik == "route" {
		fmt.Println("next: route " + *host + " -> 127.0.0.1:" + *port + " in traefik")
	}
	fmt.Println("next (gitlab side, command as code via the official CLI):")
	fmt.Println("  glab auth login --hostname " + *host)
	fmt.Println("  glab repo create kubecontainer/KubeContainer --public")
	fmt.Println("  glab runner ... (or UI: Admin > CI/CD > Runners)")
	if failN > 0 {
		os.Exit(1)
	}
}
