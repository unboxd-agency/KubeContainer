// RecordGraph — the corpus as a graph, the DBpedia pattern applied to
// the record: the prose stays canonical, the graph is extracted from
// it, never hand-written. Nodes are the corpus files; an edge is a
// reference one file makes to another. The tool emits the generated
// graph (eval/graph.txt) and returns a verdict: a reference to a path
// that does not exist is a broken edge, and broken edges fail the
// gate. It changes nothing else and acts only when invoked.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var topDirs = map[string]bool{
	"api": true, "cmd": true, "config": true, "deploy": true,
	"docs": true, "eval": true, "hack": true, "internal": true,
	"registry": true, "release": true, "site": true, "test": true,
	".github": true, ".githooks": true,
}

// Generated at build time; their absence in the tree is not a broken edge.
var generated = []string{"dist/", "bin/"}

var spanRe = regexp.MustCompile("`([^`\n]+)`")
var jsonRe = regexp.MustCompile(`[A-Za-z0-9_.-]+(?:/[A-Za-z0-9_.-]+)+`)
var bareRe = regexp.MustCompile(`[A-Za-z0-9_-]+\.[A-Za-z0-9]+`)
var lineRef = regexp.MustCompile(`:\d+$`)
var urlRe = regexp.MustCompile(`https?://[^\s)>"'` + "`" + `]+`)

func main() {
	planned := plannedRefs()
	nodes := trackedFiles()

	edges := map[string]bool{}
	var broken []string
	for _, n := range nodes {
		// Generated artifacts are not authored references: the graph
		// files cite every node by construction, and the bound books
		// only restate prose that lives canonically elsewhere.
		if n == "eval/graph.txt" || n == "eval/graph.jsonld" ||
			strings.HasPrefix(n, "site/book/") || strings.HasPrefix(n, "site-autonomyx/books/") ||
			strings.HasPrefix(n, "site-autonomyx/whitepapers/") {
			continue
		}
		raw, err := os.ReadFile(n)
		if err != nil {
			continue
		}
		for _, t := range tokens(string(raw), !strings.HasSuffix(n, ".md")) {
			t = relative(n, t)
			if t == n || !keep(t) {
				continue
			}
			if isGenerated(t) || planned[t] {
				continue
			}
			if _, err := os.Stat(t); err != nil {
				broken = append(broken, fmt.Sprintf("%s -> %s", n, t))
				continue
			}
			edges[fmt.Sprintf("%s -> %s", n, t)] = true
		}
		// External links: every URL the record cites is an edge to the
		// world outside the tree. Recorded, not liveness-checked — the
		// desk has no network; the living-anchor test is the field's.
		for _, u := range urlRe.FindAllString(string(raw), -1) {
			u = strings.TrimRight(u, ".,;:")
			if u != "" {
				edges[fmt.Sprintf("%s -> %s", n, u)] = true
			}
		}
	}

	lines := make([]string, 0, len(edges))
	for e := range edges {
		lines = append(lines, e)
	}
	sort.Strings(lines)
	out := "# RecordGraph — generated from the corpus; do not edit. Rebuild: make graph\n" +
		fmt.Sprintf("# nodes=%d edges=%d\n", len(nodes), len(lines)) +
		strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile("eval/graph.txt", []byte(out), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := writeJSONLD(nodes, lines); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("record graph: %d nodes, %d edges -> eval/graph.txt, eval/graph.jsonld\n", len(nodes), len(lines))
	if len(broken) > 0 {
		for _, b := range broken {
			fmt.Printf("[broken] %s\n", b)
		}
		fmt.Printf("verdict: %d broken reference(s) — the record does not resolve\n", len(broken))
		os.Exit(1)
	}
	fmt.Println("verdict: every reference resolves")
}

// tokens extracts path-like references. In markdown only backticked
// spans count (prose mentions a path by quoting it); everywhere else
// (code, manifests, scripts) any path-shaped string counts.
func tokens(content string, raw bool) []string {
	var spans []string
	if raw {
		// A match preceded by '/' is a tail of an absolute or URL path
		// (Kubernetes REST paths like /api/v1/...), not a file reference.
		for _, idx := range jsonRe.FindAllStringIndex(content, -1) {
			if idx[0] > 0 && content[idx[0]-1] == '/' {
				continue
			}
			spans = append(spans, content[idx[0]:idx[1]])
		}
		// Bare names (monitor.yaml) are candidates too: they resolve
		// against the source's own directory, or not at all.
		for _, idx := range bareRe.FindAllStringIndex(content, -1) {
			if idx[0] > 0 && strings.ContainsRune("/.-_", rune(content[idx[0]-1])) {
				continue
			}
			spans = append(spans, content[idx[0]:idx[1]])
		}
	} else {
		for _, m := range spanRe.FindAllStringSubmatch(content, -1) {
			// A span may list several files, comma-separated; later
			// bare names inherit the directory of the first.
			dir := ""
			for part := range strings.SplitSeq(m[1], ",") {
				part = strings.TrimSpace(part)
				if !strings.Contains(part, "/") && dir != "" {
					part = dir + part
				}
				spans = append(spans, part)
				if i := strings.LastIndex(part, "/"); i >= 0 {
					dir = part[:i+1]
				}
			}
		}
	}
	var out []string
	for _, s := range spans {
		s = lineRef.ReplaceAllString(s, "")
		s = strings.TrimSuffix(s, "/")
		if s == "" || strings.ContainsAny(s, " *${}()§\"'") {
			continue
		}
		out = append(out, s)
	}
	return out
}

// keep admits only resolved path references: rooted in a directory of
// the record. Bare names that found no neighbor are dropped silently —
// a word with a dot in it is not a claim about a file.
func keep(t string) bool {
	return strings.Contains(t, "/") && topDirs[strings.SplitN(t, "/", 2)[0]]
}

// relative resolves a bare-name reference (no slash) against the
// referencing file's own directory — the way kustomization resources
// name their neighbors. Only an existing neighbor counts; anything
// else passes through unchanged.
func relative(src, t string) string {
	if strings.Contains(t, "/") || !strings.Contains(t, ".") {
		return t
	}
	if i := strings.LastIndex(src, "/"); i >= 0 {
		neighbor := src[:i+1] + t
		if _, err := os.Stat(neighbor); err == nil {
			return neighbor
		}
	}
	return t
}

// trackedFiles lists every file of the record: the tree as git keeps
// it — every file a node, nothing generated, nothing untracked.
func trackedFiles() []string {
	raw, err := exec.Command("git", "ls-files").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var out []string
	for l := range strings.SplitSeq(strings.TrimSpace(string(raw)), "\n") {
		if l != "" {
			out = append(out, l)
		}
	}
	sort.Strings(out)
	return out
}

// writeJSONLD renders the same graph as schema.org JSON-LD: every
// corpus file a CreativeWork, every edge a citation — standard linked
// data, the DBpedia family's interchange format.
func writeJSONLD(nodes, edges []string) error {
	cites := map[string][]map[string]string{}
	all := map[string]bool{}
	for _, n := range nodes {
		all[n] = true
	}
	for _, e := range edges {
		parts := strings.SplitN(e, " -> ", 2)
		cites[parts[0]] = append(cites[parts[0]], map[string]string{"@id": parts[1]})
		all[parts[1]] = true
	}
	ids := make([]string, 0, len(all))
	for id := range all {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	graph := make([]map[string]any, 0, len(ids))
	for _, n := range ids {
		entry := map[string]any{"@id": n, "@type": "CreativeWork", "name": n}
		if c := cites[n]; len(c) > 0 {
			entry["citation"] = c
		}
		graph = append(graph, entry)
	}
	doc := map[string]any{"@context": "https://schema.org", "@graph": graph}
	raw, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("eval/graph.jsonld", append(raw, '\n'), 0o644)
}

// plannedRefs reads eval/graph-planned.txt: declared-future paths the
// roadmap names before they exist. Shrink-only, like the vocabulary
// baseline — a planned path that lands leaves the file.
func plannedRefs() map[string]bool {
	out := map[string]bool{}
	raw, err := os.ReadFile("eval/graph-planned.txt")
	if err != nil {
		return out
	}
	for l := range strings.SplitSeq(string(raw), "\n") {
		l = strings.TrimSpace(l)
		if l != "" && !strings.HasPrefix(l, "#") {
			out[l] = true
		}
	}
	return out
}

func isGenerated(t string) bool {
	for _, g := range generated {
		if strings.HasPrefix(t, g) {
			return true
		}
	}
	return false
}
